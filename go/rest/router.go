package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sijirama/x/go/rest/handlers"
)

func generateApp() *fiber.App{
    app := fiber.New()

    // create healthcheck router
    app.Get("/health" , func(c *fiber.Ctx)error {
       return c.SendString("OK") 
    })

    // create the library group and routes.
    libGroup := app.Group("/library")
    libGroup.Get("/" , handlers.TestHandler)
    libGroup.Post("/create" , handlers.CreateLibrary)

    return app
}
