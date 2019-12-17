package _63_isUgly

/*
	编写一个程序判断给定的数是否为丑数。

	丑数就是只包含质因数 2, 3, 5 的正整数。

	示例 1:

	输入: 6
	输出: true
	解释: 6 = 2 × 3
	示例 2:

	输入: 8
	输出: true
	解释: 8 = 2 × 2 × 2
	示例 3:

	输入: 14
	输出: false
	解释: 14 不是丑数，因为它包含了另外一个质因数 7。
	说明：

	1 是丑数。
	输入不会超过 32 位有符号整数的范围: [−1 << 31,  1 << 31]。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/ugly-number
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	for num != 1 {
		if num % 5 == 0 {
			num /= 5
			continue
		}
		if num % 3 == 0 {
			num /= 3
			continue
		}
		if num % 2 == 0 {
			num /= 2
			continue
		}
		return false
	}
	return true
}