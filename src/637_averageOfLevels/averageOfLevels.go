package _37_averageOfLevels

/*
	给定一个非空二叉树, 返回一个由每层节点平均值组成的数组.

	示例 1:

	输入:
		3
	   / \
	  9  20
		/  \
	   15   7
	输出: [3, 14.5, 11]
	解释:
	第0层的平均值是 3,  第1层是 14.5, 第2层是 11. 因此返回 [3, 14.5, 11].
	注意：

	节点值的范围在32位有符号整数范围内。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/average-of-levels-in-binary-tree
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {

}

func VLR(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := root.Left
	right

}