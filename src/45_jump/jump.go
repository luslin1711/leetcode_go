package _5_jump

/*
	给定一个非负整数数组，你最初位于数组的第一个位置。
	数组中的每个元素代表你在该位置可以跳跃的最大长度。
	你的目标是使用最少的跳跃次数到达数组的最后一个位置。
	示例:
	输入: [2,3,1,1,4]
	输出: 2
	解释: 跳到最后一个位置的最小跳跃数是 2。
	     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
	说明:
	假设你总是可以到达数组的最后一个位置。
	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/jump-game-ii
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func jump(nums []int) int {
	var (
		i, length int
		ans, maxPos, right int
	)
	length = len(nums)
	if length <= 1 {
		return 0
	}
	for i=0; i < length-1; i++ {
		if nums[i] + i > maxPos {
			maxPos = nums[i] + i
		}
		if i == right {
			ans++
			right = maxPos
		}
	}
	return ans
}

func Jump(nums []int) int {
	return jump(nums)
}