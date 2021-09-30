package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/allegro/bigcache"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/dgraph-io/ristretto"
	"github.com/go-redis/redis"
	"github.com/patrickmn/go-cache"
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

func WriteGoCache() {
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set("foo", "bar", cache.DefaultExpiration)
}

func ReadGoCache() {
	c := cache.New(5*time.Minute, 10*time.Minute)
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}
}

func WriteRistretto() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}

	// set a value with a cost of 1
	cache.Set("key", "value", 1)
}

func ReadRistretto() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}

	// set a value with a cost of 1
	cache.Set("key", "value", 1)

	// wait for value to pass through buffers
	cache.Wait()

	_, found := cache.Get("key")
	if !found {
		panic("missing value")
	}
}

func WriteBigcache() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	cache.Set("my-unique-key", []byte("value"))
}

func ReadBigcache() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	cache.Set("my-unique-key", []byte("value"))

	entry, _ := cache.Get("my-unique-key")
	fmt.Println(string(entry))
}

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	fmt.Println("This package is used for caches testing")
}
