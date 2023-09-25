package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/itaypanda/fiber-auth-example/database"
	"github.com/itaypanda/fiber-auth-example/router"
	"github.com/itaypanda/fiber-auth-example/utils"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	// Load environment variables
  err = godotenv.Load()
  if err != nil {
    log.Fatalf("Error loading .env file %v", err)
  }

	// Migrate & connect to the database
	database.Connect()

	// Create the web server
	app := fiber.New()

	// Setup the public routes (non-authenticated & authenticated)
	router.SetupPublicRotues(app)

	// Register middlware
	secret, err := utils.GetJWTSecret()
	if err != nil {
		log.Fatalf("Error getting JWT secret: %v", err)
	}
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
	}))

	// Setup the private routes (authenticated only)
	router.SetupPrivateRotues(app)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
