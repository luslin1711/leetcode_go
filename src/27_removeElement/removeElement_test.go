package _7_removeElement

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{0,1,2,2,3,0,4,2}
	fmt.Println(RemoveElement(nums,2))
	fmt.Println(nums)
}