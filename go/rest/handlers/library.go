package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sijirama/x/go/rest/database"
	"github.com/sijirama/x/go/rest/models"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson"
)

type libraryDTO struct {
	Name    string   `json:"name" bson:"name"`
	Address string   `json:"address" bson:"address"`
	Empty   []string `json:"no_exists" bson:"books"`
}

func DeleteLibrary(c *fiber.Ctx) error {
	docId := c.Params("id")
	libraryCollection := database.GetCollection(*database.DBclient, "libraries")
    filter := bson.M{"_id": docId}
	res, err := libraryCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		panic(err)
	}
	return c.SendString(fmt.Sprint("Document has been deleted." , res.DeletedCount , " was deleted"))
}

// GET: Get all libraries route handler
func GetLibraries(c *fiber.Ctx) error {
	libraryCollection := database.GetCollection(*database.DBclient, "libraries")
	cursor, error := libraryCollection.Find(context.TODO(), bson.M{})
	if error != nil {
		return error
	}

	var libraries []models.Library
	if err := cursor.All(context.TODO(), &libraries); err != nil {
		return err
	}

	return c.JSON(libraries)
}

// POST: Create Library route handler
func CreateLibrary(c *fiber.Ctx) error {
	nLibrary := new(libraryDTO)

	if err := c.BodyParser(nLibrary); err != nil {
		return err
	}

	nLibrary.Empty = make([]string, 0)

	libraryCollection := database.GetCollection(*database.DBclient, "libraries")
	nDoc, error := libraryCollection.InsertOne(context.TODO(), nLibrary)
	if error != nil {
		return error
	}

	return c.JSON(fiber.Map{"id": nDoc.InsertedID})
}
