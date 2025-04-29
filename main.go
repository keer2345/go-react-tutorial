package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        int    `json:"id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

// var collection *mongo.Collection

func main() {
	// Initialize a new Fiber app
	// app := fiber.New()

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
	// 设置连接超时时间（如需要）
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// 连接到 MongoDB
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		panic(err)
	}

	// 检查连接是否成功
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")

	// 使用完成后断开连接
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// PORT := os.Getenv("PORT")

	// todos := []Todo{}

	// Start the server on port 3000
	// log.Fatal(app.Listen(":" + PORT))
}
