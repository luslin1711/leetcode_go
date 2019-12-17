package _03_NumArray

import (
	"fmt"
	"testing"
)

func TestNumArray_SumRange(t *testing.T) {
	array := Constructor([]int{-2,0,3,-5,2,-1})
	fmt.Println(array.SumRange(0,2))
}

func TestNumArray_SumRange1(t *testing.T) {
	array := Constructor([]int{-2,0,3,-5,2,-1})
	fmt.Println(array.SumRange(2,5))
}

func TestNumArray_SumRange2(t *testing.T) {
	array := Constructor([]int{-1})
	fmt.Println(array.SumRange(0,0))
}
