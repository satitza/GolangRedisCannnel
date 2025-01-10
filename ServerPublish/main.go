package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {

	fmt.Println("Redis server publish")

	// Start a Redis connection
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer conn.Close()

	channel := "example-channel"
	// Publish messages periodically
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Hello, Redis! Message #%d", i)
		_, err := conn.Do("PUBLISH", channel, message)
		if err != nil {
			log.Printf("Failed to publish message to channel %s: %v", channel, err)
			continue
		}
		fmt.Printf("Published message to channel %s: %s\n", channel, message)
		time.Sleep(2 * time.Second)
	}

}
