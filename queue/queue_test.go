package queue

import (
	"testing"
)

func BenchmarkUnsafePut(b *testing.B) {
	var Q = NewUnsafeQueue(uint32(b.N + 2))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Q.Put(3)
	}
}

func BenchmarkUnsafeGet(b *testing.B) {
	var Q = NewUnsafeQueue(uint32(b.N))
	for i := 0; i < b.N; i++ {
		Q.Put(3)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Q.Get()
	}
}
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func BenchmarkUnsafe16PutGet(b *testing.B) {

	var Q = NewUnsafeQueue16(65000)
	b.ResetTimer()
	for i, cb := 0, Min(b.N, 64000); i < b.N; i, cb = i+64000, Min(b.N-i, 64000) {
		for j := 0; j < cb; j++ {
			Q.Put(3)
		}
		for j := 0; j < cb; j++ {
			Q.Get()
		}
	}
}

func BenchmarkNolockPut(b *testing.B) {
	var Q = NewQueue(uint32(b.N))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Q.Put(3)
	}
}

func BenchmarkNolockGet(b *testing.B) {
	var Q = NewQueue(uint32(b.N))
	for i := 0; i < b.N; i++ {
		Q.Put(3)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Q.Get()
	}
}
