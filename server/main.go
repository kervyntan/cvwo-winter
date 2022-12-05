package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello")
	// var myName string = "Test"
	// var myName2 = "Test2"
	// myName3 := "Test3"

	// // Formatted String
	// fmt.Printf("My life is %v and my life is %v, I like %v", myName, myName2, myName3)
	// Create an instance of the app
	app := fiber.New()

	// Get request with a callback function
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("THIS IS A TEST")
	})

	// listen to port 4000 (need the :)
	app.Listen(":4000")

	log.Fatal(app.Listen(":4000"))
}
