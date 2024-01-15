package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sijirama/x/go/rest/database"
	"github.com/sijirama/x/go/rest/models"
	"go.mongodb.org/mongo-driver/bson"
)

type bookDTO struct {
	Author    string `json:"author" bson:"author"`
	Title     string `json:"title" bson:"title"`
	ISBN      string `json:"isbn" bson:"isbn"`
	LibraryId string `json:"libraryId" bson:"libraryId"`
}

func CreateBook(c *fiber.Ctx) error {
	nBook := new(bookDTO)
	if err := c.BodyParser(nBook); err != nil {
		return err
	}
	fmt.Println(nBook)

	// get the collection
	Coll := database.GetCollection(*database.DBclient, "libraries")

	//filter
	filter := bson.D{{Key: "_id", Value: nBook.LibraryId}}
    fmt.Println(filter)

	nBookData := models.Book{
		Author: nBook.Author,
		Title:  nBook.Title,
		ISBN:   nBook.ISBN,
	}

    updatePayload := bson.D{{Key:"$push", Value: bson.M{"books": nBookData}}}

	//update the library
	res, err := Coll.UpdateOne(context.TODO(), filter, updatePayload)
	if err != nil {
		return nil
	}

    fmt.Println(res.UpsertedID)
	return c.SendString("Book has been created successfully")
}
