package _59_buddyStrings


/*
	给定两个由小写字母构成的字符串 A 和 B ，只要我们可以通过交换 A 中的两个字母得到与 B 相等的结果，就返回 true ；否则返回 false 。

	示例 1：

	输入： A = "ab", B = "ba"
	输出： true
	示例 2：

	输入： A = "ab", B = "ab"
	输出： false
	示例 3:

	输入： A = "aa", B = "aa"
	输出： true
	示例 4：

	输入： A = "aaaaaaabc", B = "aaaaaaacb"
	输出： true
	示例 5：

	输入： A = "", B = "aa"
	输出： false
	 

	提示：

	0 <= A.length <= 20000
	0 <= B.length <= 20000
	A 和 B 仅由小写字母构成。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/buddy-strings
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func buddyStrings(A string, B string) bool {
	var (
		a_char, b_char byte
	)
	a_len := len(A)
	b_len := len(B)
	if a_len != b_len {
		return false
	}
	flag := false
	end_flag := false

	for i := 0; i < a_len; i++ {

		if A[i] != B[i] {
			if flag {
				if end_flag {
					return false
				}
				end_flag = true
				if b_char != A[i] || a_char != B[i] {
					return false
				}
			} else {
				a_char = A[i]
				b_char = B[i]
				flag = true
			}
		}
	}
	if !flag {
		return false
	}
	return true
}