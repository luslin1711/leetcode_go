package _04_sumOfLeftLeaves



/*
	计算给定二叉树的所有左叶子之和。

	示例：

		3
	   / \
	  9  20
		/  \
	   15   7

	在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/sum-of-left-leaves
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if  root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return sumOfLeftLeaves(root.Right) + root.Left.Val
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func SumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeaves(root)
}