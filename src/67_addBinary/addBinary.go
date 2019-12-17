package _7_addBinary

import (
	"bytes"
	"strings"
)

/*
	给定两个二进制字符串，返回他们的和（用二进制表示）。

	输入为非空字符串且只包含数字 1 和 0。

	示例 1:

	输入: a = "11", b = "1"
	输出: "100"
	示例 2:

	输入: a = "1010", b = "1011"
	输出: "10101"
	在真实的面试中遇到过

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/add-binary
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func addBinary(a string, b string) string {
	var (
		a_len, b_len,max_len int
		buff *bytes.Buffer
		str []byte
		bit byte
	)
	a_len = len(a)
	b_len = len(b)
	if a_len > b_len {
		max_len = a_len
		b = strings.Repeat("0",a_len-b_len) + b
	} else {
		max_len = b_len
		a = strings.Repeat("0",b_len-a_len) + a
	}
	str = make([]byte,max_len+1,max_len+1)
	bit = '0'
	for i:= max_len-1; i >=0; i-- {
		sum := bit + a[i] + b[i]
		switch sum {
		case 144:
			str[i+1] = '0'
			bit = '0'
		case 145:
			str[i+1] = '1'
			bit = '0'
		case 146:
			str[i+1] = '0'
			bit = '1'
		case 147:
			str[i+1] = '1'
			bit = '1'
		}
	}
	if bit == '1' {
		str[0] = '1'
	} else {
		str = str[1:]
	}
	buff = bytes.NewBuffer(str)
	return buff.String()
}

func AddBinary(a string, b string) string {
	return addBinary(a,b)
}