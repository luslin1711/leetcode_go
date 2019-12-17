package _45_reverseVowels


/*
	编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

	示例 1:

	输入: "hello"
	输出: "holle"
	示例 2:

	输入: "leetcode"
	输出: "leotcede"
	说明:
	元音字母不包含字母"y"。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/reverse-vowels-of-a-string
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func reverseVowels(s string) string {
	if len(s) <= 1 {
		return s
	}
	chars := []byte(s)
	i, j := 0, len(chars)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}