package main

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	data  map[int]string
	cache Cacher
}

func NewStore(c Cacher) *Store {
	data := map[int]string{
		1: "let's learn golang hands on!",
		2: "foo is not bar and bar is not baz",
		3: "let's play dota2",
	}

	return &Store{
		data:  data,
		cache: c,
	}
}

// func (s *Store) getFromCache(key int) (string, bool) {
// 	val, ok := s.cache.Get(key)
// 	if ok {
// 		fmt.Println("returning key from cache")
// 		return val, ok
// 	}

// 	return "", false
// }

func (s *Store) Get(key int) (string, error) {
	val, ok := s.cache.Get(key)
	if ok {
		// busting the cache
		if err := s.cache.Remove(key); err != nil {
			fmt.Println(err)
		}
		fmt.Println("returning key from the cache")
		return val, nil
	}

	val, ok = s.data[key]
	if !ok {
		return "", fmt.Errorf("key not found: %d", key)
	}

	if err := s.cache.Set(key, val); err != nil {
		return "", err
	}

	fmt.Println("returning key from the storage")

	return val, nil
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	s := NewStore(NewRedisCache(client))
	for q := 0; q < 10; q++ {
		val, err := s.Get(1)
		if err != nil {
			log.Fatal()
		}

		fmt.Println(val)
	}

}
