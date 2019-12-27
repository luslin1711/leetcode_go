package _24_pivotIndex

import (
	"fmt"
	"testing"
)


func TestPivotIndex(t *testing.T) {
	fmt.Print(pivotIndex([]int{1,7,3,6,5,6}))
}

func TestPivotIndex2(t *testing.T) {
	fmt.Print(pivotIndex([]int{-1,-1,-1,-1,-1,0}))
}