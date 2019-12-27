package _21_shortestToChar

import "strings"

/*
	给定一个字符串 S 和一个字符 C。返回一个代表字符串 S 中每个字符到字符串 S 中的字符 C 的最短距离的数组。

	示例 1:

	输入: S = "loveleetcode", C = 'e'
	输出: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
	说明:

	字符串 S 的长度范围为 [1, 10000]。
	C 是一个单字符，且保证是字符串 S 里的字符。
	S 和 C 中的所有字母均为小写字母。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/shortest-distance-to-a-character
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func shortestToChar(S string, C byte) []int {
	length := len(S)
	ans := make([]int, length, length)
	slow := -20001
	quick := strings.Index(S,string(C))
	for i := 0; i < length; i++ {
		if S[i] != C {
			ans[i] = min(i-slow, quick-i)
		} else {
			// 找到下一个quick位置
			if i == length -1 {
				if S[i] == C {
					slow = quick
					quick = i
				} else {
					quick = 20000
				}
			} else {
				slow = quick
				quick = strings.Index(S[i+1:],string(C))
				if quick == -1 {
					quick = 20000
				} else {
					quick += i + 1
				}
			}
		}
	}
	return ans
}


func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}