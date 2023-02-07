package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/allegro/bigcache"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/dgraph-io/ristretto"
	"github.com/ducminhgd/gao/generator"
	"github.com/go-redis/redis"
	"github.com/patrickmn/go-cache"
)

var ctx = context.Background()

var MemcacheClient *memcache.Client = memcache.New("127.0.0.1:11211")
var rdb = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var kdb = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6380",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var dfdb = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6381",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func WriteRedis() {
	uuid := generator.NewUUID()
	err := rdb.Set(uuid, uuid, 60*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func ReadRedis() {
	uuid := generator.NewUUID()
	rdb.Set(uuid, uuid, 60*time.Second).Err()
	_, err := rdb.Get(uuid).Result()
	if err != nil {
		panic(err)
	}
}

func WriteKeyDb() {
	uuid := generator.NewUUID()
	err := kdb.Set(uuid, uuid, 60*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func ReadKeyDb() {
	uuid := generator.NewUUID()
	kdb.Set(uuid, uuid, 60*time.Second).Err()
	_, err := kdb.Get(uuid).Result()
	if err != nil {
		panic(err)
	}
}

func WriteDragonFlyDB() {
	uuid := generator.NewUUID()
	err := dfdb.Set(uuid, uuid, 60*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func ReadDragonFlyDB() {
	uuid := generator.NewUUID()
	dfdb.Set(uuid, uuid, 60*time.Second).Err()
	_, err := dfdb.Get(uuid).Result()
	if err != nil {
		panic(err)
	}
}

func WriteMemcache() {
	uuid := generator.NewUUID()
	MemcacheClient.Set(&memcache.Item{Key: uuid, Value: []byte(uuid), Expiration: 600})
}

func ReadMemcache() {
	uuid := generator.NewUUID()
	MemcacheClient.Set(&memcache.Item{Key: uuid, Value: []byte(uuid), Expiration: 600})
	MemcacheClient.Get(uuid)
}

func WriteGoCache() {
	uuid := generator.NewUUID()
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set(uuid, uuid, cache.DefaultExpiration)
}

func ReadGoCache() {
	uuid := generator.NewUUID()
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set(uuid, uuid, cache.DefaultExpiration)
	foo, found := c.Get(uuid)
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

	uuid := generator.NewUUID()
	// set a value with a cost of 1
	cache.Set(uuid, uuid, 1)
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
