package _42_isPowerOfFour

/*
	给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。

	示例 1:

	输入: 16
	输出: true
	示例 2:

	输入: 5
	输出: false

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/power-of-four
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func isPowerOfFour(num int) bool {
	if num <= 0  {
		return false
	}
	count := 0
	i := 0
	index := 0
	for num != 0 {
		if num & 1 == 1 {
			count++
			index = i
			if count > 1 {
				return false
			}
		}
		num = num >> 1
		i++
	}
	return index % 2 == 0
}

func IsPowerOfFour(num int) bool {
	return isPowerOfFour(num)
}