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

package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/robert-cronin/jueju/backend/internal/database"
	"github.com/robert-cronin/jueju/backend/internal/models"
)

func RequestPoem(c *fiber.Ctx) error {
	var input models.PoemRequestDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	userID := c.Locals("userID").(uuid.UUID)

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user"})
	}

	// Check if credits need to be reset
	if time.Since(user.LastCreditReset) >= 30*24*time.Hour {
		user.PoemCredits = 10
		user.LastCreditReset = time.Now()
		database.DB.Save(&user)
	}

	if user.PoemCredits <= 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":             "Insufficient credits",
			"credits_required":  1,
			"credits_available": user.PoemCredits,
		})
	}

	poemRequest := models.Poem{
		UserID:       userID,
		Prompt:       input.Prompt,
		Status:       "pending",
		AttemptCount: 1,
	}

	if err := database.DB.Create(&poemRequest).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create poem request"})
	}

	user.PoemCredits--
	database.DB.Save(&user)

	// Create DTO for response
	responseDTO := models.PoemRequestDTO{
		ID:           poemRequest.ID,
		UserID:       poemRequest.UserID,
		Prompt:       poemRequest.Prompt,
		Status:       poemRequest.Status,
		AttemptCount: poemRequest.AttemptCount,
		CreatedAt:    poemRequest.CreatedAt,
	}

	return c.JSON(responseDTO)
}

func GetUserPoemRequests(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var poemRequests []models.Poem
	if err := database.DB.Where("user_id = ?", userID).Find(&poemRequests).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch poem requests"})
	}

	// Convert to DTOs
	var responseDTOs []models.PoemResponseDTO
	for _, pr := range poemRequests {
		responseDTOs = append(responseDTOs, models.PoemResponseDTO{
			ID:        pr.ID,
			UserID:    pr.UserID,
			Prompt:    pr.Prompt,
			Poem:      pr.Poem,
			Status:    pr.Status,
			CreatedAt: pr.CreatedAt,
			UpdatedAt: pr.UpdatedAt,
		})
	}

	return c.JSON(responseDTOs)
}
