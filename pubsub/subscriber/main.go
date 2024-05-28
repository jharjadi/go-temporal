package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"os"
	"time"
)

const (
	testSub      = "testSub"
	testSubTopic = "testSubTopic"
	resultTopic  = "resultTopic"
	resultSub    = "resultSub"
)

func main() {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")

	ctx, _ := context.WithTimeout(context.Background(), 10000*time.Second)
	app, err := newApp(Config{context: ctx, gcpProjectName: "test",
		subscriptionName: testSub, topicName: resultTopic,
		options: []option.ClientOption{option.WithoutAuthentication()}})
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	app.run()

}

type App struct {
	context context.Context
	client  *pubsub.Client
	config  Config
}

type Config struct {
	context          context.Context
	gcpProjectName   string
	subscriptionName string
	topicName        string
	options          []option.ClientOption
}

func newApp(config Config) (*App, error) {
	client, err := pubsub.NewClient(config.context, config.gcpProjectName, config.options...)
	if err != nil {
		return nil, err
	}

	return &App{context: config.context, client: client, config: config}, nil
}

func (app *App) run() {
	log.Println("waiting for messages on ")
	err := app.client.Subscription(app.config.subscriptionName).
		Receive(app.config.context, func(ctx context.Context, message *pubsub.Message) {
			var messageJson map[string]interface{}
			json.Unmarshal(message.Data, &messageJson)
			log.Printf("received message with id: %s and content %v", message.ID, messageJson)
			messageJson["processed_time"] = time.Now()
			result, _ := json.Marshal(messageJson)
			app.client.Topic(app.config.topicName).Publish(ctx, &pubsub.Message{Data: result})
			message.Ack()
		})

	if err != nil {
		fmt.Println(err)
	}
	log.Println("stopped waiting for messages")
}
