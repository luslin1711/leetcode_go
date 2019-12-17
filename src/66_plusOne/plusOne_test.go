package _6_plusOne

import (
	"fmt"
	"testing"
)

func TestPlusOne1(t *testing.T) {
	fmt.Println(PlusOne([]int{1,2,3}))
}

func TestPlusOne2(t *testing.T) {
	fmt.Println(PlusOne([]int{9,9}))
}

func TestPlusOne3(t *testing.T) {
	fmt.Println(PlusOne([]int{9}))
}

func TestPlusOne4(t *testing.T) {
	fmt.Println(PlusOne([]int{0}))
}