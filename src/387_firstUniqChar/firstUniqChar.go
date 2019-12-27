package _87_firstUniqChar

import "math"

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.
 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func firstUniqChar(s string) int {
	related := make([]int,26,26)
	length := len(s)
	for i := 0; i < length; i++ {
		if related[s[i]-'a'] == 0 {
			related[s[i]-'a'] = i + 1
		} else {
			related[s[i]-'a'] -= length
		}
	}
	minIndex := math.MaxInt32
	flag := false
	for i := 0; i < len(related); i++ {
		if related[i] > 0 {
			flag = true
			if minIndex > related[i] {
				minIndex = related[i]
			}
		}
	}
	if minIndex >= 1 && flag {
		return minIndex-1
	}
	return -1
}

func FirstUniqChar(s string) int {
	return firstUniqChar(s)
}