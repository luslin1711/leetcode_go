package _05_sortArrayByParity



/*
	给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。

	你可以返回满足此条件的任何数组作为答案。

	 

	示例：

	输入：[3,1,2,4]
	输出：[2,4,3,1]
	输出 [4,2,3,1]，[2,4,1,3] 和 [4,2,1,3] 也会被接受。
	 

	提示：

	1 <= A.length <= 5000
	0 <= A[i] <= 5000

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/sort-array-by-parity
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func sortArrayByParity(A []int) []int {
	length := len(A)
	if length <= 1 {
		return A
	}
	p := 0
	q := length -1
	for p < q {
		for p < q && A[p] % 2 == 0 {
			p++
		}
		for p < q && A[q] % 2 == 1 {
			q--
		}
		A[p], A[q] = A[q], A[p]
	}
	return A
}