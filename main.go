package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	data map[int]string
}

func NewStore() *Store {
	data := map[int]string{
		1: "let's learn golang hands on!",
		2: "foo is not bar and bar is not baz",
		3: "let's play dota2",
	}

	return &Store{
		data: data,
	}
}

func (s *Store) Get(key int) (string, error) {
	val, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("key not found: %d", key)
	}

	return val, nil
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

}
