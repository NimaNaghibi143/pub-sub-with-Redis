package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ch := make(chan any)

	ch <- 4
	ch <- "hello"
	ch <- struct{}{} // 0 bytes
	ch <- errors.New("error")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	sub := client.Subscribe(ctx, "coords")

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", msg.Payload)
	}
}
