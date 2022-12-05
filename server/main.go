package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Println("Hello")

	todos := []Todo{}
	// Create an instance of the app
	app := fiber.New()

	// Get request with a callback function
	// First endpoint
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("THIS IS A TEST")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		// define a variable to catch the error (err)
		// Then test the error, if it's not nil (it is an actual error), then return the error
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.JSON(todos)
	})

	// listen to port 4000 (need the :)
	app.Listen(":4000")

	log.Fatal(app.Listen(":4000"))
}
