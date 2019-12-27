package _41_reverseStr

import "bytes"

/*
	给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。如果剩余少于 k 个字符，则将剩余的所有全部反转。如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。

	示例:

	输入: s = "abcdefg", k = 2
	输出: "bacdfeg"
	要求:

	该字符串只包含小写的英文字母。
	给定字符串的长度和 k 在[1, 10000]范围内。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/reverse-string-ii
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func reverseStr(s string, k int) string {
	if k <= 1 {
		return s
	}
	length := len(s)
	left_num := length % (2 * k)
	ans := []byte(s)
	s = s[:length-left_num]
	for i := 0; i * k < length - left_num; i++ {
		if i % 2 == 0 {
			reserve(ans,i*k, (i+1)*k-1)
		}
	}
	if left_num >= k {
		reserve(ans,length-left_num,length-left_num+k-1)
	} else {
		reserve(ans,length-left_num,length-1)
	}
	return bytes.NewBuffer(ans).String()
}

func reserve(chars []byte,start, end int)  {
	p := start
	q := end
	for p < q {
		chars[p], chars[q] = chars[q], chars[p]
		p++
		q--
	}
}

func ReverseStr(s string, k int) string {
	return reverseStr(s, k)
}