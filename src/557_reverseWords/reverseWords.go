package _57_reverseWords

import "bytes"

/*
	给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

	示例 1:

	输入: "Let's take LeetCode contest"
	输出: "s'teL ekat edoCteeL tsetnoc" 
	注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func reverseWords(s string) string {
	chars := []byte(s)
	i := 0
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		break
	}
	slow,quick := i,i
	for ; quick <= len(s); quick++ {
		if quick == len(s) {
			reserve(chars,slow,quick-1)
			break
		}
		if chars[quick] == ' ' {
			reserve(chars,slow,quick-1)
			slow = quick + 1
		}
	}
	return bytes.NewBuffer(chars).String()

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

func ReverseWords(s string) string {
	return reverseWords(s)
}