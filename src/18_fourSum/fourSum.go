package _8_fourSum

/*
	给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
	注意：
	答案中不可以包含重复的四元组。
	示例：
	给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
	满足要求的四元组集合为：
	[
	  [-1,  0, 0, 1],
	  [-2, -1, 1, 2],
	  [-2,  0, 0, 2]
	]
	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/4sum
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func fourSum(nums []int, target int) [][]int {
	var (
		i,j int
		start,end,sum int
		last_start, last_end int
		length int
		result [][]int
	)
	length = len(nums)
	if length < 4 {
		return result
	}
	QuickSort(nums,0, length-1)
	result = make([][]int,0, length / 4)
	for i=0; i < length-3;i++ {
		if target < nums[i] + nums[i+1] + nums[i+2] + nums[i+3] {
			break
		}
		if target > nums[i] + nums[length-1] + nums[length-2] + nums[length-3] {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j=i+1; j<length-2;j++ {
			if target < nums[i] + nums[j] + nums[j+1] + nums[j+2] {
				continue
			}
			if target > nums[i] + nums[j] + nums[length-1] + nums[length-2] {
				continue
			}
			if j>i+1 && nums[j] == nums[j-1] {
				continue
			}
			start = j+1
			end = length-1
			for start < end {
				sum = nums[i] + nums[j] + nums[start] + nums[end]
				if sum < target {
					start++
				} else if sum > target {
					end--
				} else {
					result = append(result, []int{nums[i],nums[j],nums[start],nums[end]})
					last_start = start
					last_end = end
					for start < end && nums[start] == nums[last_start] {
						start++
					}
					for start < end && nums[end] == nums[last_end] {
						end--
					}
				}
			}
		}
	}
	return result
}
func FourSum(nums []int, target int) [][]int {
	return fourSum(nums, target)
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