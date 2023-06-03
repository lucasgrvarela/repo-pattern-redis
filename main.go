package main

import (
	"context"
	"fmt"
)

func main() {
	redisClient := NewRedisClient("localhost:6379")

	key := "mykey"

	err := redisClient.Set(context.Background(), key, "myvalue")
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
	fmt.Println("Key set successfully!")

	val, err := redisClient.Get(context.Background(), key)
	if err != nil {
		fmt.Println("Error getting value:", err)
		return
	}
	fmt.Printf("Key: '%s' retrieved, value: '%s'\n", key, val)
}
