package main

import (
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"

	// db stuff
	"github.com/joho/godotenv"
	"github.com/sijirama/x/go/rest/database"
)

func main() {

	initError, client := InitApp()

	if initError != nil {
		log.Fatal(initError)
		panic(initError)
	}

	defer database.CloseMongoDb(client)

	app := generateApp()

	// get port from env
	PORT := os.Getenv("PORT")
	app.Listen(":" + PORT)
}

func InitApp() (error, *mongo.Client) {
	envError := LoadEnv()
	if envError != nil {
		return errors.New("Error with Environment varaibales"), nil
	}

	dbError, client := database.StartMongoDb()
	if dbError != nil {
		return errors.New("Error with database"), nil
	}

	return nil, client
}

func LoadEnv() error {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return errors.New("No .env file found")
	}
	return nil
}
