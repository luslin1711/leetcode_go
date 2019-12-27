package _28_maximumProduct

import (
	"fmt"
	"testing"
)


func TestMaximumProduct(t *testing.T) {
	fmt.Println(MaximumProduct([]int{1,2,3}))
}

func TestMaximumProduct2(t *testing.T) {
	fmt.Println(MaximumProduct([]int{1,2,3,4}))
}

func TestMaximumProduct3(t *testing.T) {
	fmt.Println(MaximumProduct([]int{-11,-2,3,4}))
}