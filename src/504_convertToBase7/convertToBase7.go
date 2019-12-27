package _04_convertToBase7

import "bytes"

/*
	给定一个整数，将其转化为7进制，并以字符串形式输出。

	示例 1:

	输入: 100
	输出: "202"
	示例 2:

	输入: -7
	输出: "-10"
	注意: 输入范围是 [-1e7, 1e7] 。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/base-7
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func convertToBase7(num int) string {
	buff := bytes.NewBuffer(nil)
	if num < 0 {
		buff.WriteByte('-')
		num = -num
	} else if num == 0 {
		return "0"
	}
	chars := make([]byte,0)
	for num != 0 {
		chars = append(chars,byte(num % 7 + '0'))
		num /= 7
	}
	reverse(chars)
	buff.Write(chars)
	return buff.String()
}

func reverse(chars []byte)  {
	p := 0
	q := len(chars) - 1
	for p < q {
		chars[p], chars[q] = chars[q],chars[p]
		p++
		q--
	}
}

func ConvertToBase7(num int) string {
	return convertToBase7(num)
}