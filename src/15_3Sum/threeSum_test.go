package _5_3Sum

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1,23,321,54,1,25,1,23,4,213,4,12,3,4,1234,21354,1,341}
	QuickSort(arr,0,len(arr)-1)
	fmt.Println(arr)
}

func TestThreeSum(t *testing.T) {
	arr := []int{-1,0,1,2,-1,4}
	fmt.Println(ThreeSum(arr))
}