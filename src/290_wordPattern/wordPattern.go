package _90_wordPattern

import (
	"strings"
)

/*
	给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

	这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

	示例1:

	输入: pattern = "abba", str = "dog cat cat dog"
	输出: true
	示例 2:

	输入:pattern = "abba", str = "dog cat cat fish"
	输出: false
	示例 3:

	输入: pattern = "aaaa", str = "dog cat cat dog"
	输出: false
	示例 4:

	输入: pattern = "abba", str = "dog dog dog dog"
	输出: false
	说明:
	你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。    

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/word-pattern
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func wordPattern(pattern string, str string) bool {
	pToSmap := make(map[byte]string)
	sToPmap := make(map[string]byte)
	strs := strings.Split(str, " ")
	if len(pattern) != len(strs) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		s, exist := pToSmap[pattern[i]]
		if !exist {
			pToSmap[pattern[i]] = strs[i]
		} else {
			if s != strs[i] {
				return false
			}
		}
		p, exist := sToPmap[strs[i]]
		if !exist {
			sToPmap[strs[i]] = pattern[i]
		} else {
			if   p != pattern[i]{
				return false
			}
		}
	}
	return true
}

func WordPattern(pattern string, str string) bool {
	return wordPattern(pattern,str)
}
