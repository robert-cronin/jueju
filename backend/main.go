package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/joho/godotenv"

	"github.com/robert-cronin/jueju-backend/api"
	_ "github.com/robert-cronin/jueju-backend/api"
)

func middleware(c *fiber.Ctx) error {
	fmt.Println("Request to", c.Path())
	return c.Next()
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found, resorting to the environment")
	}

	server := api.NewServer()

	app := fiber.New()
	app.Use(cors.New())

	apiGroup := app.Group("/api", middleware)

	api.RegisterHandlers(apiGroup, server)

	log.Fatal(app.Listen("0.0.0.0:3000"))

}
