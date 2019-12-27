package _15_addStrings

import (
	"bytes"
	"strings"
)

/*
	给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

	注意：

	num1 和num2 的长度都小于 5100.
	num1 和num2 都只包含数字 0-9.
	num1 和num2 都不包含任何前导零。
	你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/add-strings
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func addStrings(num1 string, num2 string) string {
	n1_len := len(num1)
	n2_len := len(num2)
	if n1_len == 0 {
		return num2
	}
	if n2_len == 0 {
		return num1
	}
	max_len := n2_len
	if n1_len > n2_len {
		max_len = n1_len
		num2 = strings.Repeat("0",n1_len-n2_len) + num2
	} else {
		num1 = strings.Repeat("0",n2_len-n1_len) + num1
	}
	strs := make([]byte,max_len+1,max_len+1)
	bit := byte(0)
	for i := max_len-1; i >=0; i-- {
		sum := bit + num1[i] + num2[i] - 96
		if sum >= 10 {
			strs[i+1] = sum - 10 + 48
			bit = 1
		} else  {
			strs[i+1] = sum + 48
			bit = 0
		}
	}
	if bit == 1 {
		strs[0] = '1'
	} else {
		strs = strs[1:]
	}
	buff := bytes.NewBuffer(strs)
	return buff.String()
}


func AddStrings(num1 string, num2 string) string {
	return addStrings(num1,num2)
}