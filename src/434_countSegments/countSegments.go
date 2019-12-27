package _34_countSegments



/*
	统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。

	请注意，你可以假定字符串里不包括任何不可打印的字符。

	示例:

	输入: "Hello, my name is John"
	输出: 5

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/number-of-segments-in-a-string
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	nums := 0
	flag := false
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && !flag {
			continue
		} else if s[i] == ' ' && flag && s[i-1] != ' '{
			nums++
		}
		flag = true
	}
	if s[len(s)-1] != ' '{
		nums++
	}
	return nums
}

func CountSegments(s string) int {
	return countSegments(s)
}