package _32_findPairs

import (
	"math"
	"sort"
)

/*
	给定一个整数数组和一个整数 k, 你需要在数组里找到不同的 k-diff 数对。这里将 k-diff 数对定义为一个整数对 (i, j), 其中 i 和 j 都是数组中的数字，且两数之差的绝对值是 k.

	示例 1:

	输入: [3, 1, 4, 1, 5], k = 2
	输出: 2
	解释: 数组中有两个 2-diff 数对, (1, 3) 和 (3, 5)。
	尽管数组中有两个1，但我们只应返回不同的数对的数量。
	示例 2:

	输入:[1, 2, 3, 4, 5], k = 1
	输出: 4
	解释: 数组中有四个 1-diff 数对, (1, 2), (2, 3), (3, 4) 和 (4, 5)。
	示例 3:

	输入: [1, 3, 1, 5, 4], k = 0
	输出: 1
	解释: 数组中只有一个 0-diff 数对，(1, 1)。
	注意:

	数对 (i, j) 和数对 (j, i) 被算作同一数对。
	数组的长度不超过10,000。
	所有输入的整数的范围在 [-1e7, 1e7]。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/k-diff-pairs-in-an-array
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func findPairs(nums []int, k int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}
	sort.Ints(nums)
	if nums[length-1] - nums[0] < k {
		return 0
	} else if nums[length-1] - nums[0] == k{
		return 1
	}

	ans := 0
	pre := nums[length-1]

	for i := 0; i < length; i++ {
		if nums[i] == pre {
			continue
		}
		pre = nums[i]
		qpre := math.MaxInt64
		for q := i+1; q < len(nums);q++ {
			if nums[q] == qpre {
				continue
			}
			qpre = nums[q]
			if nums[q] - nums[i] == k {
				ans++
			} else if nums[q] - nums[i] > k {
					break
			}
		}
	}
	return ans
}


func FindPairs(nums []int, k int) int {
	return findPairs(nums, k)
}

