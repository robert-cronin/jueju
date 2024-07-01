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

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Auth0ID         string    `gorm:"uniqueIndex;not null" json:"auth0_id"`
	Email           string    `gorm:"uniqueIndex;not null" json:"email"`
	EmailVerified   bool      `json:"email_verified"`
	Name            string    `json:"name"`
	Nickname        string    `json:"nickname"`
	Picture         string    `json:"picture"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	LastLogin       time.Time `json:"last_login"`
	PoemCredits     int       `gorm:"default:10" json:"poem_credits"`
	LastCreditReset time.Time `json:"last_credit_reset"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	u.LastCreditReset = time.Now()
	return nil
}
