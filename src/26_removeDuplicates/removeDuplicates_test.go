package _6_removeDuplicates

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1,1,2}
	fmt.Println(RemoveDuplicates(nums))
	fmt.Println(nums)
}