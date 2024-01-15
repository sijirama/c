package main

import (
	//"context"
	"errors"
	"log"

	//"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	// db stuff
	"github.com/joho/godotenv"
	//"go.mongodb.org/mongo-driver/bson"
	"github.com/sijirama/x/go/rest/database"
)

func main() {

	initError, client := InitApp()

	if initError != nil {
		log.Fatal(initError)
		panic(initError)
	}

	defer database.CloseMongoDb(client)

	//app := fiber.New()
    app := generateApp()


	// app.Post("/", func(c *fiber.Ctx) error {
 //        // write a todo to the database
	//
 //        sameple := bson.M{"name":"sample todo"}    
 //        coll := database.GetCollection( *client, "todos")
	//
 //        result , err := coll.InsertOne(context.TODO() , sameple)
 //        if err != nil {
 //            return c.Status(fiber.StatusInternalServerError).SendString("Error inserting Todo!")
 //        }
	//
 //        return c.JSON(result)
	// })

	errr := app.Listen(":3000")
	if errr != nil {
		log.Panic("Server Error: ", errr)
	}
	log.Println("Server listening on Port :", 3000)
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
