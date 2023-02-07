package main

import (
	"testing"
)

// func BenchmarkWriteMemcache(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		WriteMemcache()
// 	}
// }

// func BenchmarkReadMemcache(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		ReadMemcache()
// 	}
// }

// func BenchmarkWriteMemcacheGoroutine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		go WriteMemcache()
// 	}
// }

// func BenchmarkReadMemcacheGoroutine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		go ReadMemcache()
// 	}
// }

func BenchmarkWriteRedis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteRedis()
	}
}

func BenchmarkReadRedis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadRedis()
	}
}

// func BenchmarkWriteRedisGoroutine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		go WriteRedis()
// 	}
// }

// func BenchmarkReadRedisGoroutine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		go ReadRedis()
// 	}
// }

// func BenchmarkWriteKeyDb(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		WriteKeyDb()
// 	}
// }

// func BenchmarkReadKeyDb(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		ReadKeyDb()
// 	}
// }

// func BenchmarkWriteKeyDbGoroutine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		go WriteKeyDb()
// 	}
// }

// func BenchmarkReadKeyDbGoroutine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		go ReadKeyDb()
// 	}
// }

func BenchmarkWriteDragonFlyDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteDragonFlyDB()
	}
}

func BenchmarkReadDragonFlyDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadDragonFlyDB()
	}
}

// func BenchmarkWriteDragonFlyDBGoroutine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		go WriteDragonFlyDB()
// 	}
// }

// func BenchmarkReadDragonFlyDBGoroutine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		go ReadDragonFlyDB()
// 	}
// }

// func BenchmarkWriteGoCache(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		WriteGoCache()
// 	}
// }

// func BenchmarkReadGoCache(b *testing.B) {
// 	c := cache.New(5*time.Minute, 10*time.Minute)
// 	c.Set("foo", "bar", cache.DefaultExpiration)
// 	for i := 0; i < b.N; i++ {
// 		c.Get("foo")
// 	}
// }

// func BenchmarkWriteGoCacheGoroutine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		go WriteGoCache()
// 	}
// }

// func BenchmarkReadGoCacheGoRoutine(b *testing.B) {
// 	c := cache.New(5*time.Minute, 10*time.Minute)
// 	c.Set("foo", "bar", cache.DefaultExpiration)
// 	for i := 0; i < b.N; i++ {
// 		go c.Get("foo")
// 	}
// }

// func BenchmarkWriteRistretto(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		WriteRistretto()
// 	}
// }

// func BenchmarkReadRistretto(b *testing.B) {
// 	cache, err := ristretto.NewCache(&ristretto.Config{
// 		NumCounters: 1e7,     // number of keys to track frequency of (10M).
// 		MaxCost:     1 << 30, // maximum cost of cache (1GB).
// 		BufferItems: 64,      // number of keys per Get buffer.
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	// set a value with a cost of 1
// 	cache.Set("key", "value", 1)

// 	// wait for value to pass through buffers
// 	cache.Wait()
// 	for i := 0; i < b.N; i++ {
// 		cache.Get("key")
// 	}
// }

// func BenchmarkWriteRistrettoGoroutine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		go WriteRistretto()
// 	}
// }

// func BenchmarkReadRistrettoGoroutine(b *testing.B) {
// 	cache, err := ristretto.NewCache(&ristretto.Config{
// 		NumCounters: 1e7,     // number of keys to track frequency of (10M).
// 		MaxCost:     1 << 30, // maximum cost of cache (1GB).
// 		BufferItems: 64,      // number of keys per Get buffer.
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	// set a value with a cost of 1
// 	cache.Set("key", "value", 1)

// 	// wait for value to pass through buffers
// 	cache.Wait()
// 	for i := 0; i < b.N; i++ {
// 		go cache.Get("key")
// 	}
// }

// func BenchmarkWriteBigcache(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		WriteBigcache()
// 	}
// }

// func BenchmarkReadBigcache(b *testing.B) {
// 	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
// 	cache.Set("my-unique-key", []byte("value"))
// 	for i := 0; i < b.N; i++ {
// 		cache.Get("my-unique-key")
// 	}
// }

// func BenchmarkWriteBigcacheGoroutine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		go WriteBigcache()
// 	}
// }

// func BenchmarkReadBigcacheGoroutine(b *testing.B) {
// 	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
// 	cache.Set("my-unique-key", []byte("value"))
// 	for i := 0; i < b.N; i++ {
// 		go cache.Get("my-unique-key")
// 	}
// }
