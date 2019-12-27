package _97_findShortestSubArray

import (
	"fmt"
	"testing"
)


func TestFindShortestSubArray(t *testing.T) {
	fmt.Println(findShortestSubArray([]int{1,2,2,3,1}))
}



func TestFindShortestSubArray2(t *testing.T) {
	fmt.Println(findShortestSubArray([]int{1,2,2,3,1,1,4,2}))
}

func TestFindShortestSubArray3(t *testing.T) {
	fmt.Println(findShortestSubArray([]int{2,1,1,2,1,3,3,3,1,3,1,3,2}))
}