package _06_findRelativeRanks

import (
	"fmt"
	"testing"
)


func TestQuickSort(t *testing.T) {
	nums := []int{1,5,3,7,2,88,54,321,6,0,4315,134,5,7,8}
	fmt.Println(QuickSort(nums))
	fmt.Println(nums)

}

func TestFindRelativeRanks(t *testing.T) {
	nums := []int{1,5,3,7,2,88,54,321,6,0,4315,134,5,7,8}
	fmt.Println(FindRelativeRanks(nums))
}