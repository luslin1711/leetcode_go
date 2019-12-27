package _05_sortArrayByParity

import (
	"fmt"
	"testing"
)


func TestSortArrayByParity(t *testing.T) {
	fmt.Println(sortArrayByParity([]int{3,1,2,4}))
}

func TestSortArrayByParity1(t *testing.T) {
	fmt.Println(sortArrayByParity([]int{1,3,1,2,4,1}))
}