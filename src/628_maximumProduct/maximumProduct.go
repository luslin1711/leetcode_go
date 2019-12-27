package _28_maximumProduct


/*
	给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

	示例 1:

	输入: [1,2,3]
	输出: 6
	示例 2:

	输入: [1,2,3,4]
	输出: 24
	注意:

	给定的整型数组长度范围是[3,104]，数组中所有的元素范围是[-1000, 1000]。
	输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/maximum-product-of-three-numbers
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func maximumProduct(nums []int) int {
	max1 , max2 , max3 := -1001, -1001, -1001
	min1, min2 := 1001,1001
	for i := 0; i < len(nums); i++ {
		if nums[i] > max1 {
			max2,max3 = max1, max2
			max1 = nums[i]
		} else if nums[i] > max2 {
			max3 = max2
			max2 = nums[i]
		} else if nums[i] > max3 {
			max3 = nums[i]
		}
		if nums[i] < min1 {
			min2 = min1
			min1 = nums[i]
		} else if nums[i] < min2 {
			min2 = nums[i]
		}
	}
	return max(max1 * max2 * max3, max1 * min1 * min2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaximumProduct(nums []int) int {
	return maximumProduct(nums)
}