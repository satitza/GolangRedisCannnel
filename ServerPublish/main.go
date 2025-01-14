package main

import (
	"encoding/json"
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

	channel := "channel-notifications"
	// Publish messages periodically
	for i := 1; i <= 1; i++ {

		allUserId := []string{"c371bd39-ac80-4ca0-b285-51e0eb0a96e6"}

		byteArray, err := json.Marshal(allUserId)
		if err != nil {
			log.Fatalf("Error marshaling to byte array: %v", err)
		}

		_, err = conn.Do("PUBLISH", channel, byteArray)
		if err != nil {
			log.Printf("Failed to publish message to channel %s: %v", channel, err)
			continue
		}
		fmt.Printf("Published message to channel %s: %v\n", channel, allUserId)
		time.Sleep(2 * time.Second)
	}

}
