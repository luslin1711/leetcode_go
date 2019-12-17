package _25_isPalindrome

/*
	给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

	说明：本题中，我们将空字符串定义为有效的回文串。

	示例 1:

	输入: "A man, a plan, a canal: Panama"
	输出: true
	示例 2:

	输入: "race a car"
	输出: false

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/valid-palindrome
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func isPalindrome(s string) bool {
	var (
		i,j,length int
	)
	length = len(s)
	j = length-1
	i=0
	for i < j {
		if !inRange(s[i]) {
			i++
			continue
		}
		if !inRange(s[j]){
			j--
			continue
		}
		if !same(s[i],s[j]){
			return false
		}
		i++
		j--
	}
	return true
}

func inRange(a byte) bool {
	if  (a >= 97 && a <= 122) || (a >= 65 && a<=90) || (a >= 48 && a<=57)   {
		return true
	}
	return false
}
func same(a,b byte) bool {
	if a == b {
		return true
	}
	if a >57 && b > 57 {
		if a - 32 == b {
			return true
		}
		if b - 32 == a {
			return true
		}
	}
	return false
}

func IsPalindrome(s string) bool {
	return isPalindrome(s)
}