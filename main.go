package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"url-shortener-go/config"
	"url-shortener-go/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDB()
	config.MigrateDB(db)

	r := fiber.New()

	r.Post("/api/shorten", handlers.ShortenURL(db))
	r.Get("/:shortId", handlers.ResolveURL(db))

	log.Fatal(r.Listen(":8080"))
}