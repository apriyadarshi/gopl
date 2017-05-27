package popcount_test

import (
	"gopl/ch2/popcount"
	"testing"
)

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount2(255255255)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount3(255255255)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount4(255255255)
	}
}
