package _31_isPowerOfTwo

import (
	"fmt"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {

	for i :=0 ; i < 20; i++ {
		fmt.Println(i,IsPowerOfTwo(i))
	}
}