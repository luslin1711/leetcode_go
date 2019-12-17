package _2_IntToRoman

import (
	"testing"
)

const num = 1994

func BenchmarkIntToRoman(b *testing.B) {
	for i := 0; i <b.N; i++ {
		IntToRoman(num)
	}
}

func BenchmarkIntToRoman2(b *testing.B) {
	for i := 0; i <b.N; i++ {
		IntToRoman2(num)
	}
}

