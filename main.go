package main

import (
	"context"
	"fmt"
)

func main() {
	redisClient := NewRedisClient("localhost:6379")

	err := redisClient.Set(context.Background(), "mykey", "myvalue")
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
	fmt.Println("Key set successfully!")
}
