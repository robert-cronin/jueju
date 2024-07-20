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

type Poem struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User         User      `gorm:"foreignKey:UserID" json:"-"`
	Prompt       string    `gorm:"type:text;not null" json:"prompt"`
	Poem         string    `gorm:"type:text" json:"poem"`
	Status       string    `gorm:"type:varchar(20);not null" json:"status"`
	AttemptCount int       `gorm:"default:0" json:"attempt_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (pr *Poem) BeforeCreate(tx *gorm.DB) error {
	pr.ID = uuid.New()
	pr.Status = "pending"
	pr.AttemptCount = 0
	return nil
}
