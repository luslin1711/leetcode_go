package _45_findErrorNums

import (
	"fmt"
	"testing"
)


func TestFindErrorNums(t *testing.T) {
	fmt.Println(FindErrorNums([]int{1,2,2,4}))
}


func TestFindErrorNums2(t *testing.T) {
	fmt.Println(FindErrorNums([]int{1,1}))
}