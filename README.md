# Create project

```sh
mkdir go-react-tutorial
cd go-react-tutorial

go mod init github.com/keer2345/go-react-tutorial
```

### fiber

⚡️ Express inspired web framework written in Go

```sh
go get -u github.com/gofiber/fiber/v3

```

### air-verse/air

☁️ Air - Live reload for Go apps

```sh
go install github.com/air-verse/air@latest # install air
air init  # initialize air
air  # run project
```

## VS Code Extensions

- Even Better TOML
- Thunder Client (~~REST Client~~ | ~~Postman~~)

# Add Todo item

```go

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

```

Run:

```sh
> curl -X POST -H "Content-Type: application/json" --data "{\"body\":\"hello world\"}" localhost:3000/api/todo
{"id":1,"completed":false,"body":"hello world"}
```
