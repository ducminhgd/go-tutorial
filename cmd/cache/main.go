package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis"
)

var ctx = context.Background()

var MemcacheClient *memcache.Client = memcache.New("10.0.0.1:11211")
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var kdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6380",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func WriteRedis() {
	err := rdb.Set("key", "value", 60*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func ReadRedis() {
	_, err := rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
}

func WriteKeyDb() {
	err := kdb.Set("key", "value", 60*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func ReadKeyDb() {
	_, err := kdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
}

func WriteMemcache() {
	MemcacheClient.Set(&memcache.Item{Key: "my-key", Value: []byte("my value"), Expiration: 600})
}

func ReadMemcache() {
	MemcacheClient.Get("my-key")
}

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	fmt.Println("This package is used for memcached testing")
}
