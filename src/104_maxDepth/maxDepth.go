package _04_maxDepth


/*
	给定一个二叉树，找出其最大深度。

	二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

	说明: 叶子节点是指没有子节点的节点。

	示例：
	给定二叉树 [3,9,20,null,null,15,7]，

		3
	   / \
	  9  20
		/  \
	   15   7
	返回它的最大深度 3 。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */



type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func maxDepth(root *TreeNode) int {
	var (
		left_n, right_n int
	)
	if root == nil {
		return 0
	} else {
		left_n = maxDepth(root.Left)
		right_n = maxDepth(root.Right)
		if left_n > right_n {
			return left_n+1
		} else {
			return right_n+1
		}
	}
}


func MaxDepth(root *TreeNode) int {
	return maxDepth(root)
}