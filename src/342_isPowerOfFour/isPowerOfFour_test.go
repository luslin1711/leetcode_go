package _42_isPowerOfFour

import (
	"fmt"
	"testing"
)

func TestIsPowerOfFour(t *testing.T) {
	fmt.Println(IsPowerOfFour(16))
}

func TestIsPowerOfFour1(t *testing.T) {
	fmt.Println(IsPowerOfFour(64))
}

func TestIsPowerOfFour2(t *testing.T) {
	fmt.Println(IsPowerOfFour(1))
}

func TestIsPowerOfFour3(t *testing.T) {
	fmt.Println(IsPowerOfFour(0))
}