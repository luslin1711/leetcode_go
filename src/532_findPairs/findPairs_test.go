package _32_findPairs

import (
	"fmt"
	"testing"
)

func TestFindPairs(t *testing.T) {
	fmt.Println(FindPairs([]int{3,1,4,1,5},2))
}


func TestFindPairs2(t *testing.T) {
	fmt.Println(FindPairs([]int{3,3,1,4,1,5},2))
}


func TestFindPairs3(t *testing.T) {
	fmt.Println(FindPairs([]int{1,2,3,4,5},1))
}
