package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/option"
	"os"
	"time"
)

//const (
//	testSub      = "in-sub"
//	testSubTopic = "in-topic"
//	resultTopic  = "out-topic"
//	resultSub    = "out-sub"
//)

func main() {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	app, err := newApp(Config{context: ctx, gcpProjectName: "test",
		subscriptionName: testSub, topicName: resultTopic,
		options: []option.ClientOption{option.WithoutAuthentication()}})

	if err != nil {
		fmt.Errorf("Error: %s", err)

	}
	app.client.Subscription(resultSub).Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var jsonMessage map[string]interface{}
		json.Unmarshal(msg.Data, &jsonMessage)
		fmt.Println(jsonMessage["greeting"])
		//assert.Equal(t, "hello", jsonMessage["greeting"], "greeting field is kept as is")
		//assert.NotEmpty(t, jsonMessage["processed_time"], "processed time field is added")
		//fmt.Println("finished assertions")
	})
}
