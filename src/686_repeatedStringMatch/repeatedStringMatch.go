package _86_repeatedStringMatch

import "strings"

/*
	给定两个字符串 A 和 B, 寻找重复叠加字符串A的最小次数，使得字符串B成为叠加后的字符串A的子串，如果不存在则返回 -1。

	举个例子，A = "abcd"，B = "cdabcdab"。

	答案为 3， 因为 A 重复叠加三遍后为 “abcdabcdabcd”，此时 B 是其子串；A 重复叠加两遍后为"abcdabcd"，B 并不是其子串。

	注意:

	 A 与 B 字符串的长度在1和10000区间范围内。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/repeated-string-match
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func repeatedStringMatch(A string, B string) int {
	endl := len(B) + 2 * len(A)
	for i :=1 ; len(A) * i <= endl; i++ {
		str := strings.Repeat(A,i)
		if strings.Index(str,B) != -1 {
			return i
		}
	}
	return -1
}