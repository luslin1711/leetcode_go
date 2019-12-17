package _005_largestSumAfterKNegations

import (
	"fmt"
	"testing"
)

func TestLargestSumAfterKNegations(t *testing.T) {
	fmt.Println(LargestSumAfterKNegations([]int{4,2,3},1))
}

func TestLargestSumAfterKNegations1(t *testing.T) {
	fmt.Println(LargestSumAfterKNegations([]int{3,-1,0,2},3))
}

func TestLargestSumAfterKNegations2(t *testing.T) {
	fmt.Println(LargestSumAfterKNegations([]int{2,-3,-1,5,-4},2))
}