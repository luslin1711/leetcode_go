package _11_minDepth



/*
	给定一个二叉树，找出其最小深度。

	最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

	说明: 叶子节点是指没有子节点的节点。

	示例:

	给定二叉树 [3,9,20,null,null,15,7],

		3
	   / \
	  9  20
		/  \
	   15   7
	返回它的最小深度  2.

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left_dep := minDepth(root.Left)
	right_dep := minDepth(root.Right)
	if root.Left == nil || root.Right == nil {
		return left_dep + right_dep +1
	} else {
		if left_dep < right_dep {
			return left_dep+1
		} else {
			return right_dep+1
		}
	}
}

func MinDepth(root *TreeNode) int {
	return minDepth(root)
}