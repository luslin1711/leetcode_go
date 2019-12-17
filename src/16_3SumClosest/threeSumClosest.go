package _6_3SumClosest

import "math"

/*
	给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

	例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

	与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/3sum-closest
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func threeSumClosest(nums []int, target int) int {
	var (
		i,j,k,length,closest,sum int

	)
	length = len(nums)
	if length < 3 {
		return 0
	}
	QuickSort(nums, 0, length-1)
	closest = nums[0] + nums[1] + nums[2]
	for i=0; i < length-2;i++ {
		j = i + 1
		k = length - 1
		for j != k {
			sum = nums[i] + nums[j] + nums[k]
			if math.Abs(float64(sum-target)) < math.Abs(float64(closest-target)){
				closest = sum
			}
			if  sum > target {
				k--
				for j != k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < target {
				j++
				for j != k && nums[j] == nums[j-1] {
					j++
				}
			} else {
				return target
			}
		}
		for i < length-2 && nums[i] == nums[i+1]{
			i++
		}
	}
	return closest
}

func ThreeSumClosest(nums []int, target int) int {
	return threeSumClosest(nums,target)
}

func point(arr []int, left, right int) int {
	var (
		start, end int
	)
	start = left
	end = right
	tmpInt := arr[left]
	for start < end {
		for ;start < end && arr[end] >= tmpInt ; end--{}
		for ;start < end && arr[start] <= tmpInt; start++{}

		arr[start], arr[end] = arr[end], arr[start]
	}
	arr[start] , arr[left] = arr[left], arr[start]
	return start
}

func QuickSort(arr []int, left, right int)  {
	if left < right {
		i := point(arr, left, right)
		QuickSort(arr, left, i -1)
		QuickSort(arr, i+1, right)
	}
}