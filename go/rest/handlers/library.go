package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sijirama/x/go/rest/database"
)

type libraryDTO struct {
    Name string `json:"name" bson:"name"`
    Address string `json:"address" bson:"address"`
}

func CreateLibrary(c *fiber.Ctx) error{
    nLibrary := new(libraryDTO)

    if err := c.BodyParser(nLibrary); err != nil {
        return err
    }

    libraryCollection := database.GetCollection(*database.DBclient, "libraries")
    nDoc , error := libraryCollection.InsertOne(context.TODO() , nLibrary)
    if error != nil {
        return error
    }

    return c.JSON(fiber.Map{"id":nDoc.InsertedID})
}
