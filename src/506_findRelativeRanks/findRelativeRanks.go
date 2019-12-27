package _06_findRelativeRanks

import "strconv"

func point(nums, related_nums []int, left, right int) int {

	tmp := nums[left]
	l := left
	r := right
	for l != r {
		for nums[r] >= tmp && l < r{
			r--
		}
		for nums[l] <= tmp && l < r {
			l++
		}
		nums[l],nums[r] = nums[r], nums[l]
		related_nums[l],related_nums[r] = related_nums[r],related_nums[l]
	}
	nums[l], nums[left] = nums[left], nums[l]
	related_nums[l],related_nums[left] = related_nums[left],related_nums[l]

	return l
}

func quickSort(nums, related_nums []int,start, end int)  {
	if start >= end {
		return
	}
	i := point(nums, related_nums,start,end)
	quickSort(nums, related_nums,start,i-1)
	quickSort(nums, related_nums,i+1,end)
}

func QuickSort(nums []int)  []int {
	related_nums := make([]int,len(nums),len(nums))
	for i := 0; i < len(related_nums); i++ {
		related_nums[i] = i
	}
	quickSort(nums, related_nums,0,len(nums)-1)
	return related_nums
}

//{1,5,3,7,2,88,54,321,6,0,4315,134,5,7,8}
//[9 0 4 2 1 12 8  3 13 14 6 5 11 7 10]
func findRelativeRanks(nums []int) []string {
	length := len(nums)
	ans := make([]string,length,length)
	ranks := QuickSort(nums)
	for i := 0; i < length; i++ {
		if i == length -1 {
			ans[ranks[i]] = "Gold Medal"
		} else if i == length -2 {
			ans[ranks[i]] = "Silver Medal"
		} else if i == length -3 {
			ans[ranks[i]] = "Bronze Medal"
		} else {
			ans[ranks[i]] = strconv.Itoa(length-i)
		}

	}
	return ans
}

func FindRelativeRanks(nums []int) []string {
	return findRelativeRanks(nums)
}