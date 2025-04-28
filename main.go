package main

import (
	"fmt"
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

	// Get Todos
	app.Get("/api/todo", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(todos)
	})

	// Create a Todo
	app.Post("/api/todo", func(c fiber.Ctx) error {
		// todo := new(Todo)
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

	// Update a Todo
	app.Patch("/api/todo/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = !todo.Completed
				return c.Status(fiber.StatusOK).JSON(todos[i])
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	})

	// Delete a Todo
	app.Delete("/api/todo/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
