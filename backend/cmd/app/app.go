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

package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/robert-cronin/jueju/backend/internal/api"
	"github.com/robert-cronin/jueju/backend/internal/server"
)

func middleware(c *fiber.Ctx) error {
	fmt.Println("Request to", c.Path())
	return c.Next()
}

func Bootstrap() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found, resorting to the environment")
	}

	srv := server.NewServer()

	app := fiber.New()
	apiGroup := app.Group("/api", middleware)

	app.Use(cors.New())

	api.RegisterHandlers(apiGroup, srv)

	log.Fatal(app.Listen("0.0.0.0:3000"))
}
