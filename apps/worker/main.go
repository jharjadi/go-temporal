package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	fmt.Println("Fullname=", os.Getenv("fullname"))
	fmt.Println("Email=", os.Getenv("email"))

}
