package _43_diameterOfBinaryTree



type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
var maxdepth int = 1

func diameterOfBinaryTree(root *TreeNode) int {
	depth(root)
	return maxdepth - 1
}



func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left)
	r := depth(root.Right)
	maxdepth = max(maxdepth, l + r + 1)
	return max(l,r) + 1
}

func max(a, b int) int {
	if a > b {
		return  a
	}
	return b
}

func DiameterOfBinaryTree(root *TreeNode) int {
	return diameterOfBinaryTree(root)
}