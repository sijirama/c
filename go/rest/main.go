package main

import (
	"errors"

	"log"

	"github.com/gofiber/fiber/v2"

	// db stuff
	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	"github.com/sijirama/x/go/rest/database"
)

func main() {

	envError := LoadEnv()
	if envError != nil {
	}

	dbError, _ := database.StartMongoDb()
	if dbError != nil {
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	errr := app.Listen(":3000")
	if errr != nil {
		log.Panic("Server Error: ", errr)
	}
	log.Println("Server listening on Port :", 3000)
}

func LoadEnv() error {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return errors.New("No .env file found")
	}
	return nil
}
