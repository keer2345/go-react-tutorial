package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	todos := []Todo{}

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		// return c.SendString("Hello, World 👋!")

		// Send a json response to the client
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Hello, World 👋!"})
	})

	app.Post("/api/todo", func(c fiber.Ctx) error {
		todo := &Todo{}

		if err := c.Bind().JSON(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Todo body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(fiber.StatusCreated).JSON(todo)
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
