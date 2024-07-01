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
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/robert-cronin/jueju/backend/internal/api"
	"github.com/robert-cronin/jueju/backend/internal/config"
	"github.com/robert-cronin/jueju/backend/internal/database"
	"github.com/robert-cronin/jueju/backend/internal/redis"
	"github.com/robert-cronin/jueju/backend/internal/server"
	"github.com/spf13/viper"
)

func loggingMiddleware(c *fiber.Ctx) error {
	fmt.Println("Request to", c.Path())
	return c.Next()
}

// Register gobs
func initGob() {
	gob.Register(map[string]interface{}{})
}

func Bootstrap() {
	initGob()

	// Get the environment variable
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	viper.Set("env", env)

	// Load the configuration
	config.InitConfig()

	// Connect to the database
	database.InitDB()

	// Initialize the Redis client
	redis.Init()

	// Create the server
	srv, err := server.NewServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Create the Fiber app
	app := fiber.New()

	// Set CORS middleware
	if env == "development" {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     "http://localhost:5173",
			AllowCredentials: true,
			AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
			AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		}))
	} else {
		app.Use(cors.New(cors.Config{
			AllowOrigins: "https://jueju.robertcronin.com",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
	}

	apiV1 := app.Group("/api/v1", loggingMiddleware)

	// Register the unathenticated API handler
	authGroup := apiV1.Group("/auth")
	authGroup.Get("/login", srv.Login)
	authGroup.Get("/callback", srv.Callback)

	// Create protected routes group
	protectedGroup := apiV1.Group("")
	protectedGroup.Use(srv.Auth.AuthMiddleware())

	// Register protected routes
	protectedGroup.Get("/logout", srv.Logout)
	protectedGroup.Get("/user", srv.GetUser)
	protectedGroup.Get("/poems", srv.GetUserPoemRequests)
	protectedGroup.Post("/poems", srv.RequestPoem)

	// Use the generated handlers
	api.RegisterHandlers(apiV1, srv)

	// Start the server
	log.Fatal(app.Listen("0.0.0.0:3000"))
}
