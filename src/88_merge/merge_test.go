package _8_merge

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1,2,3,0,0,0}
	nums2 := []int{2,5}
	Merge(nums1,3,nums2,2)
	fmt.Println(nums1)
}