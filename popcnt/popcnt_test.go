package popcnt

import (
	"testing"
)

var Sink interface{}

func BenchmarkPopcnt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = popcnt(uint64(i))
	}
}

func BenchmarkPopcnt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = popcnt2(uint64(i))
	}
}
