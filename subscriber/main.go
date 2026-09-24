package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	CHANNEL_NAME = "updates:test:1"
)

func main() {
	ctx := context.Background()

	// Connect to the Redis server
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // default DB
	})

	// Subscribe to the channel
	pubsub := rdb.Subscribe(ctx, CHANNEL_NAME)
	defer pubsub.Close()

	fmt.Println("Listening for messages...")
	ist, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading IST timezone:", err)
		return
	}
	currentTime := time.Now().In(ist).Format(time.RFC822)

	// Continuously wait and print received messages
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			fmt.Printf("[%s] - Error receiving message: %s\n", currentTime, err.Error())
			continue
		}

		fmt.Printf("[%s] - Received message: %s\n", currentTime, msg.Payload)
	}
}
