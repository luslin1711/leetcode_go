package _5_3Sum

/*
	给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

	注意：答案中不可以包含重复的三元组。

	例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

	满足要求的三元组集合为：
	[
	  [-1, 0, 1],
	  [-1, -1, 2]
	]

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/3sum
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func threeSum(nums []int) [][]int {
	var (
		result [][]int
		length,index,i,j,k,preValue,preIndexValue int
		twoSum int
	)
	length = len(nums)
	result = make([][]int,0,length / 3)
	if length < 3 {
		return result
	}
	QuickSort(nums,0, length-1)
	if nums[0] > 0 || nums[length-1] <0 {
		return result
	}
	preIndexValue = -1 << 31
	for index = 0; index < length; index++ {
		if nums[index] > 0 {
			break
		}
		if nums[index] == preIndexValue{
			continue
		} else {
			preIndexValue = nums[index]
		}
		i = index
		j = length-1
		preValue = -1 << 31
		for i < j {
			if nums[j] == preValue {
				continue
			}
			preValue = nums[j]
			twoSum = nums[i] + nums[j]
			if twoSum > 0 {
				for k = i+1; twoSum + nums[k] < 0 && k < j; k++ {}
			} else {
				for k = j-1; twoSum + nums[k] > 0 && k > i; k-- {}
			}
			if twoSum + nums[k] == 0 {
				result = append(result, []int{nums[i],nums[j],nums[k]})
			}
			j--
			if twoSum + nums[j] < 0 {
				break
			}

		}

	}
	return result
}
func ThreeSum(nums []int) [][]int {
	return  threeSum(nums)
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