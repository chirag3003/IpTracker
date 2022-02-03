package main

import (
	"context"
	"git.chirag.codes/chirag/ip_tracker/db"
	"git.chirag.codes/chirag/ip_tracker/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// setting up environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//connecting MongoDB
	client := db.ConnectDB()
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// setting up of server
	app := fiber.New()
	routes.Init(app)
	log.Fatal(app.Listen(":3000"))
}
