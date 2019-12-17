package _4_LongestCommonPrefix


/*
	编写一个函数来查找字符串数组中的最长公共前缀。

	如果不存在公共前缀，返回空字符串 ""。

	示例 1:

	输入: ["flower","flow","flight"]
	输出: "fl"
	示例 2:

	输入: ["dog","racecar","car"]
	输出: ""
	解释: 输入不存在公共前缀。
	说明:

	所有输入只包含小写字母 a-z 。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/longest-common-prefix
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestCommonPrefix(strs []string) string {
	var (
		length,index,comIndex,comLen int
		commonStr string
	)
	length = len(strs)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return strs[0]
	}
	commonStr = strs[0]
	for index = 1; index < length; index++ {
		comLen = len(strs[index])
		if comLen > len(commonStr) {
			comLen = len(commonStr)
		}
		for comIndex = 0; comIndex < comLen; comIndex++ {
			if strs[index][comIndex] != commonStr[comIndex] {
				break
			}
		}
		if comIndex == 0 {
			return ""
		}
		commonStr = commonStr[:comIndex]
	}
	return commonStr
}

func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}