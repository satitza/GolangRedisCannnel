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

	channel := "example-channel"
	// Publish messages periodically
	for i := 1; i <= 5; i++ {

		allUserId := []string{"820d839c-c660-4a54-be40-8b8028245e3e", "15e7997f-7f59-42c1-aaa4-cc9ee827827c", "c4301568-5467-4c3c-acfc-698d1f66c911"}

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
