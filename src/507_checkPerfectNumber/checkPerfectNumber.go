package _07_checkPerfectNumber

import "math"

/*
	对于一个 正整数，如果它和除了它自身以外的所有正因子之和相等，我们称它为“完美数”。
	给定一个 整数 n， 如果他是完美数，返回 True，否则返回 False
	示例：
	输入: 28
	输出: True
	解释: 28 = 1 + 2 + 4 + 7 + 14
	提示：
	输入的数字 n 不会超过 100,000,000. (1e8)
	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/perfect-number
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sqrt_res := int(math.Sqrt(float64(num)))
	sum := 1
	for i := 2; i <= sqrt_res; i++ {
		div_res := num % i
		if div_res == 0 {
			sum += i
			if div_res != num / i {
				sum += num / i
			}
		}

	}
	return sum == num
}

func CheckPerfectNumber(num int) bool {
	return checkPerfectNumber(num)
}