package main

import (
	"fmt"
	"github.com/jharjadi/go-temporal/app"
	"github.com/joho/godotenv"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
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

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "greeting-tasks", worker.Options{})

	w.RegisterWorkflow(app.GreetSomeone)
	w.RegisterActivity(app.GreetInSpanish)
	w.RegisterActivity(app.FarewellInSpanish)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}

}
