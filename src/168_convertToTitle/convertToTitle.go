package _68_convertToTitle

import (
	"bytes"
	"fmt"
)

/*
	给定一个正整数，返回它在 Excel 表中相对应的列名称。

	例如，

		1 -> A
		2 -> B
		3 -> C
		...
		26 -> Z
		27 -> AA
		28 -> AB
		...
	示例 1:

	输入: 1
	输出: "A"
	示例 2:

	输入: 28
	输出: "AB"
	示例 3:

	输入: 701
	输出: "ZY"

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/excel-sheet-column-title
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func convertToTitle(n int) string {
	var (
		buff *bytes.Buffer
		ans []byte
	)
	for n > 0  {
		n--
		ans = append(ans,byte(n%26+'A'))
		n = n / 26
	}
	fmt.Println(ans)
	for i:=0 ; i <len(ans) /2; i++ {
		ans[i],ans[len(ans)-i-1] = ans[len(ans)-i-1],ans[i]
	}
	buff = bytes.NewBuffer(ans)
	return buff.String()
}


func ConvertToTitle(n int) string {
	return convertToTitle(n)
}