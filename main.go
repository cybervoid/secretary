package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// notion()
	chatgpt()

	// fmt.Printf("Database user: %s, password: %s\n", dbUser, "*****")
}
