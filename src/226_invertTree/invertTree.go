package _26_invertTree


/*
	翻转一棵二叉树。

	示例：

	输入：

		 4
	   /   \
	  2     7
	 / \   / \
	1   3 6   9
	输出：

		 4
	   /   \
	  7     2
	 / \   / \
	9   6 3   1

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/invert-binary-tree
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func InvertTree(root *TreeNode) *TreeNode {
	return invertTree(root)
}