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

package database

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/robert-cronin/jueju/backend/internal/models"
)

func seedUsers(db *gorm.DB) error {
	users := []models.User{
		{
			UID:      uuid.New().String(),
			Username: "user1",
			Email:    "user1@robertcronin.com",
		},
		{
			UID:      uuid.New().String(),
			Username: "user2",
			Email:    "user2@robertcronin.com",
		},
	}

	for _, user := range users {
		// Check if user already exists
		var existingUser models.User
		result := db.Where("email = ?", user.Email).First(&existingUser)
		if result.Error == gorm.ErrRecordNotFound {
			// User doesn't exist, create new user
			if err := db.Create(&user).Error; err != nil {
				return fmt.Errorf("error creating user: %w", err)
			}
		} else if result.Error != nil {
			// Some other error occurred
			return fmt.Errorf("error checking for existing user: %w", result.Error)
		} else {
			// User already exists, skip creation
			fmt.Printf("User with email %s already exists, skipping creation\n", user.Email)
			continue
		}

		err := seedPoems(db, user)
		if err != nil {
			return err
		}
	}

	return nil
}

func seedPoems(db *gorm.DB, user models.User) error {
	poems := []models.Poem{
		{
			Title:       "Poem One",
			Content:     "Content of Poem One",
			Translation: "Translation of Poem One",
			UserID:      user.ID,
		},
		{
			Title:       "Poem Two",
			Content:     "Content of Poem Two",
			Translation: "Translation of Poem Two",
			UserID:      user.ID,
		},
	}

	for _, poem := range poems {
		if err := db.Create(&poem).Error; err != nil {
			return fmt.Errorf("error creating pem: %w for user: %s", err, user.Username)
		}
	}

	return nil
}
