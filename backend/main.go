package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/robert-cronin/jueju-backend/internal/authenticator"
	"github.com/robert-cronin/jueju-backend/internal/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found, resorting to the environment")
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Print("Server listening on port 3000")
	if err := rtr.Run(":3000"); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
