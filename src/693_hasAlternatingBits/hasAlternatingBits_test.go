package _93_hasAlternatingBits

import (
	"fmt"
	"testing"
)


func TestHasAlternatingBits(t *testing.T) {
	for i := 0 ; i < 30; i++ {
		fmt.Println(i, "->", hasAlternatingBits(i))
	}
}