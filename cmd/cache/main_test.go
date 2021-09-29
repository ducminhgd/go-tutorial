package main

import (
	"testing"
)

func BenchmarkWriteMemcache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteMemcache()
	}
}

func BenchmarkReadMemcache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadMemcache()
	}
}

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

func BenchmarkWriteKeyDb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteKeyDb()
	}
}

func BenchmarkReadKeyDb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadKeyDb()
	}
}
