package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	CHANNEL_NAME = "updates:test:1"
)

func main() {
	ctx := context.Background()

	// Create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // no password set
		DB:       0,                // default DB
	})

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter messages to send to the subscriber. Type 'exit' to quit.")

	// Loop to accept input from the terminal
	for scanner.Scan() {
		if scanner.Err() != nil {
			fmt.Println("Error reading input:", scanner.Err())
			continue
		}

		message := scanner.Text()
		if message == "exit" {
			break
		}
		ist, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			fmt.Println("Error loading IST timezone:", err)
			return
		}
		currentTime := time.Now().In(ist).Format(time.RFC822)
		// Publish the input message to the channel
		err = rdb.Publish(ctx, CHANNEL_NAME, message).Err()
		if err != nil {
			fmt.Printf("[%s] - Error publishing message:%s\n", currentTime, err.Error())
			continue
		}

		fmt.Printf("[%s] - Message published: [%s]\n", currentTime, message)
	}
}
