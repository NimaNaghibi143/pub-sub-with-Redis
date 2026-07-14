package main

import (
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ttl := time.Second * 4
	s := NewStore(NewRedisCache(client, ttl))
	for q := 0; q < 2; q++ {
		val, err := s.Get(3)
		if err != nil {
			log.Fatal()
		}

		fmt.Println(val)

		time.Sleep(3 * time.Second)
	}

}
