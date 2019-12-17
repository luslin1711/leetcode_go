package __PalindromeNumber

/*
	判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
	示例 1:
	输入: 121
	输出: true
	示例 2:
	输入: -121
	输出: false
	解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
	示例 3:
	输入: 10
	输出: false
	解释: 从右向左读, 为 01 。因此它不是一个回文数。
 */

func isPalindrome(x int) bool {
	var (
		point int
		reverse_res,value int
	)
	if x < 0 {
		return false
	} else if  x < 10 {
		return true
	}
	reverse_res = 0
	value = x
	for value != 0 {
		point = value % 10
		reverse_res = reverse_res * 10 + point
		value = value / 10
	}
	return  reverse_res == x
}

func IsPalindrome(x int) bool {
	return isPalindrome(x)
}