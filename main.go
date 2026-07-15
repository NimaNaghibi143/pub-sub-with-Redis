package main

import "fmt"

// type Transformer interface {
// 	Transform(string) string
// }

// // CAS: content addressable storage
// type CASTransformer struct{}

// func (t *CASTransformer) Transform(val string) string {
// 	return val + "Nima Naghibi"
// }

// type Server struct {
// 	tr Transformer
// }

// func (s Server) DoStuff() {
// 	f := "readfile"
// 	filepath := s.tr.Transform(f)

// 	fmt.Println(filepath)
// }

type TransformFunc func(string) string

type Server struct {
	transFunc TransformFunc
}

func (s Server) DoStuff() {
	f := "readfile"
	filepath := s.transFunc(f)

	fmt.Println(filepath)
}

func main() {

}

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/redis/go-redis/v9"
// )

// func main() {
// 	ch := make(chan any)

// 	ch <- 4
// 	ch <- "hello"
// 	ch <- struct{}{} // 0 bytes
// 	ch <- errors.New("error")

// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,
// 	})

// 	ttl := time.Second * 4
// 	s := NewStore(NewRedisCache(client, ttl))
// 	for q := 0; q < 2; q++ {
// 		val, err := s.Get(3)
// 		if err != nil {
// 			log.Fatal()
// 		}

// 		fmt.Println(val)

// 		time.Sleep(3 * time.Second)
// 	}

// 	ctx := context.Background()
// 	for i := 0; i < 100; i++ {
// 		if err := client.Publish(ctx, "coords", i).Err(); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
