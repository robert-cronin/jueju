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
	"github.com/robert-cronin/jueju/backend/internal/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initialises a new database connection.
func InitDB() {
	// Get the dsn from the environment (default to db.dsn but if that isnt set, get from the environment)
	dsn := viper.GetString("db.dsn")
	if dsn == "" {
		dsn = viper.GetString("DATABASE_URI")
	}

	// Connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the database
	err = DB.AutoMigrate(&models.User{}, &models.Poem{}, &models.Poem{})
	if err != nil {
		panic(err)
	}

	// Seed the database in development environment
	if viper.GetString("env") != "production" {
		err = seedUsers(DB)
		if err != nil {
			panic(err)
		}
	}
}
