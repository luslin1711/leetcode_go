package _62_countPrimeSetBits

import (
	"fmt"
	"testing"
)


func TestCountPrimeSetBits(t *testing.T) {
	fmt.Println(countPrimeSetBits(6,10))
}


func TestCountPrimeSetBits1(t *testing.T) {
	fmt.Println(countPrimeSetBits(10,15))
}

func TestCountPrimeSetBits2(t *testing.T) {
	fmt.Println(countPrimeSetBits(990000,1000000))
}
