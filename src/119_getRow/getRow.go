package _19_getRow

import "fmt"

/*
	给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
	在杨辉三角中，每个数是它左上方和右上方的数的和。
	示例:
	输入: 3
	输出: [1,3,3,1]

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/pascals-triangle-ii
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func getRow(rowIndex int) []int {
	var (
		dp []int
		i,j int
	)
	dp = make([]int,rowIndex+1,rowIndex+1)
	dp[0] = 1
	if rowIndex > 0 {
		dp[1] = 1
	}
	for i = 2; i <= rowIndex; i++ {
		fmt.Println(dp)
		for j=i-1; j>0;j-- {
			dp[j] = dp[j] + dp[j-1]
		}
		dp[i] = 1
	}
	return dp[:rowIndex+1]
}



func GetRow(rowIndex int) []int {
	return getRow(rowIndex)
}