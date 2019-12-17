package _3_RomanToint

import (
	"testing"
)

func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RomanToInt("III")
	}
}

func BenchmarkRomanToInt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RomanToInt2("III")
	}
}
