package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var uri, dbname, dbpass string

	// https://blog.csdn.net/weixin_62533201/article/details/146239323
	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}
	if dbname = os.Getenv("MONGODB_NAME"); dbname == "" {
		log.Fatal("You must set your 'MONGODB_NAME' environment variable.")
	}
	if dbpass = os.Getenv("MONGODB_PASS"); dbpass == "" {
		log.Fatal("You must set your 'MONGODB_PASS' environment variable.")
	}
	uri = "mongodb://" + dbname + ":" + dbpass + "@" + uri

	// Set connect address
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.Background())

	// 检查连接是否成功
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")

	// 获取数据库和集合
	database := client.Database("go-react-tutorial")
	collection = database.Collection("todos")

	// Initialize a new Fiber app
	app := fiber.New()

	// todos := []Todo{}
	app.Get("/api/todo", getTodos)
	app.Post("/api/todo", addTodo)
	app.Patch("/api/todo/:id", updateTodo)
	// app.Delete("/api/todo/:id", deleteTodo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Start the server on port 3000
	log.Fatal(app.Listen("0.0.0.0:" + port))

}

func getTodos(c fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}
	return c.JSON(todos)
}

func addTodo(c fiber.Ctx) error {
	todo := new(Todo)

	if err := c.Bind().JSON(todo); err != nil {
		return err
	}
	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func updateTodo(c fiber.Ctx) error {
	var todo Todo
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectID}
	err = collection.FindOne(context.Background(), filter).Decode(&todo)
	if err != nil {
		panic(err)
	}

	update := bson.M{"$set": bson.M{"completed": !todo.Completed}}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}
