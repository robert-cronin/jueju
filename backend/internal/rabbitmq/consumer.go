// Copyright 2024 Robert Cronin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/robert-cronin/jueju/backend/internal/database"
	"github.com/robert-cronin/jueju/backend/internal/models"
	"gorm.io/gorm"
)

func StartPoemResponseConsumer() {
	messages, err := Client.channel.Consume(
		"poem_responses",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range messages {
			var response models.PoemResponseDTO
			err := json.Unmarshal(d.Body, &response)
			if err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}

			err = processPoemResponse(response)
			if err != nil {
				log.Printf("Error processing poem response: %v", err)
			}
		}
	}()

	log.Printf("Waiting for poem responses. To exit press CTRL+C")
	<-forever
}

func processPoemResponse(response models.PoemResponseDTO) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var poemRequest models.Poem
		if err := tx.First(&poemRequest, "id = ?", response.ID).Error; err != nil {
			return err
		}

		poemRequest.Status = response.Status
		if response.Status == "completed" {
			poemRequest.Poem = response.Poem
		} else {
			poemRequest.AttemptCount++
			if poemRequest.AttemptCount >= 5 {
				poemRequest.Status = "failed"
			} else {
				// Requeue the request
				requestDTO := models.PoemRequestDTO{
					ID:        poemRequest.ID,
					UserID:    poemRequest.UserID,
					Prompt:    poemRequest.Prompt,
					CreatedAt: poemRequest.CreatedAt,
				}
				message, err := json.Marshal(requestDTO)
				if err != nil {
					return err
				}
				if err := Client.PublishMessage("poem_requests", message); err != nil {
					return err
				}
			}
		}

		return tx.Save(&poemRequest).Error
	})
}
