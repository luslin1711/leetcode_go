package _67_isPerfectSquare

import (
	"fmt"
	"testing"
)

func TestIsPerfectSquare(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(i,"->",IsPerfectSquare(i))
	}
}

