package _36_singleNumber

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	fmt.Println(SingleNumber([]int{2,2,1}))
}

func TestSingleNumber2(t *testing.T) {
	fmt.Println(SingleNumber([]int{2,5,2,1,4,1,4}))
}