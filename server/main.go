package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
	Test  bool   `json:"test"` // default value is false
}

func main() {
	fmt.Println("Hello")

	todos := []Todo{}
	// Create an instance of the app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
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

	app.Put("/api/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid Id")
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = true
				break
			}
		}

		return c.JSON(todos)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	// listen to port 4000 (need the :)
	app.Listen(":4000")

	log.Fatal(app.Listen(":4000"))
}
