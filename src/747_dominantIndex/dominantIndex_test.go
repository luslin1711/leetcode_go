package _47_dominantIndex

import (
	"fmt"
	"testing"
)


func TestDominantIndex(t *testing.T) {
	fmt.Println(dominantIndex([]int{3,6,1,0}))
}

func TestDominantIndex1(t *testing.T) {
	fmt.Println(dominantIndex([]int{1,0}))
}