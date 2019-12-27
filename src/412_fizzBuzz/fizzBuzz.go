package _12_fizzBuzz

import "strconv"

/*
	写一个程序，输出从 1 到 n 数字的字符串表示。

	1. 如果 n 是3的倍数，输出“Fizz”；

	2. 如果 n 是5的倍数，输出“Buzz”；

	3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

	示例：

	n = 15,

	返回:
	[
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz"
	]

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/fizz-buzz
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func fizzBuzz(n int) []string {
	if n == 0 {
		return []string{}
	}
	ans := make([]string,n,n)
	for i := 0; i < n; i++ {
		if (i+1) % 3 == 0 {
			if (i+1) % 5 == 0 {
				ans[i] =  "FizzBuzz"
			} else {
				ans[i] =  "Fizz"
			}
		} else if (i+1) % 5 == 0 {
			ans[i] =  "Buzz"
		} else {
			ans[i] = strconv.Itoa(i+1)
		}
	}
	return ans
}

func FizzBuzz(n int) []string {
	return fizzBuzz(n)
}