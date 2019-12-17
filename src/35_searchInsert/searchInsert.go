package _5_searchInsert


/*
	给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
	你可以假设数组中无重复元素。
	示例 1:

	输入: [1,3,5,6], 5
	输出: 2
	示例 2:

	输入: [1,3,5,6], 2
	输出: 1
	示例 3:

	输入: [1,3,5,6], 7
	输出: 4
	示例 4:

	输入: [1,3,5,6], 0
	输出: 0

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/search-insert-position
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func searchInsert(nums []int, target int) int {
	var (
		length int
		p,q,mid int
	)
	length = len(nums)
	if length == 0 || target < nums[0]{
		return 0
	}
	p = 0
	q = length - 1
	if target > nums[q] {
		return length
	}
	for q - p > 1 {
		mid = (p+q) / 2
		if nums[mid] > target {
			q = mid
		} else if nums[mid] < target {
			p = mid
		} else {
			return mid
		}
	}
	if nums[p] == target {
		return p
	}
	return q
}

func SearchInsert(nums []int, target int) int {
	return searchInsert(nums,target)
}
