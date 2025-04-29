- https://www.youtube.com/watch?v=lNd7XlXwlho

# Create GO Project

```sh
mkdir go-react-tutorial
cd go-react-tutorial

go mod init github.com/keer2345/go-react-tutorial
```

## Modules

- fiber
- godotenv
- air

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

## Config environment

```sh
go get github.com/joho/godotenv
touch .env
```

file of `.env`:

```sh
PORT=5000
```

```go
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	// ...

	log.Fatal(app.Listen(":" + PORT))
```

# Todo

## API without DB

Add todo

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

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}

```

Run:

```sh
> curl -X POST -H "Content-Type: application/json" --data "{\"body\":\"hello world\"}" localhost:3000/api/todo
{"id":1,"completed":false,"body":"hello world"}
```

Get Todo list

```go
	// Get Todos
	app.Get("/api/todo", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(todos)
	}
```

Update Todo

```go
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

```

Delete a Todo

```go

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
```

## API with DB

https://github.com/mongodb/mongo-go-driver

https://www.mongodb.com/zh-cn/docs/drivers/go/current/fundamentals/connections/connection-guide/#std-label-golang-connection-guide

```sh
go get go.mongodb.org/mongo-driver/v2/mongo
```

# Client Project

```sh
mkdir client
cd client
npm create vite@latest .
```

Choose `react` --> `Typescript`.
