package _75_findRadius

import (
	"fmt"
	"testing"
)

func TestFindRadius(t *testing.T) {
	 fmt.Println(FindRadius([]int{1,2,3},[]int{2}))
}


func TestFindRadius2(t *testing.T) {
	fmt.Println(FindRadius([]int{1,2,3,4,5,6},[]int{3,4}))
}