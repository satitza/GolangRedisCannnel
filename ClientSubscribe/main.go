package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {

	fmt.Println("Redis client subscribe")

	// Start a Redis connection
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer conn.Close()

	go subscriber()

	select {}

}

func subscriber() {
	// Connect to Redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatalf("Failed to connect to Redis in subscriber: %v", err)
	}
	defer conn.Close()

	// Create a Redis Pub/Sub connection
	psc := redis.PubSubConn{Conn: conn}

	// Subscribe to a channel
	channel := "example-channel"
	if err := psc.Subscribe(channel); err != nil {
		log.Fatalf("Failed to subscribe to channel %s: %v", channel, err)
	}
	fmt.Printf("Subscribed to channel: %s\n", channel)

	// Listen for messages
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("Received message from %s: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("Subscription event: %s to %s (count: %d)\n", v.Kind, v.Channel, v.Count)
		case error:
			log.Printf("Error while receiving message: %v\n", v)
			return
		}
	}
}
