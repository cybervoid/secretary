package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get sensitive info from environment variables
	dbUser := os.Getenv("NOTION_TOKEN")

	fmt.Printf("Database user: %s, password: %s\n", dbUser, "*****")
}
