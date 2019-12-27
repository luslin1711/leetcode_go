package _87_longestUnivaluePath

import "math"

/*
	给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。

	注意：两个节点之间的路径长度由它们之间的边数表示。

	示例 1:

	输入:

				  5
				 / \
				4   5
			   / \   \
			  1   1   5
	输出:

	2
	示例 2:

	输入:

				  1
				 / \
				4   5
			   / \   \
			  4   4   5
	输出:

	2
	注意: 给定的二叉树不超过10000个结点。 树的高度不超过1000。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/longest-univalue-path
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	 ans := 0
	 LRV(root, &ans)
	 return ans
}


func LRV(root *TreeNode, ans *int) int {
	if root == nil {
		return math.MinInt64
	}
	left := LRV(root.Left, ans)
	right := LRV(root.Right, ans)

	samel := 0
	samer := 0

	if root.Left != nil && root.Val == root.Left.Val {
		samel = left + 1
	}

	if root.Right != nil && root.Val == root.Right.Val {
		samer = right + 1
	}
	*ans = max(*ans, samer + samel)
	return max(samel, samer)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}