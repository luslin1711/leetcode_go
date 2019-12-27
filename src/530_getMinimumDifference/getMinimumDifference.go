package _30_getMinimumDifference

import "math"

/*
	给定一个所有节点为非负值的二叉搜索树，求树中任意两节点的差的绝对值的最小值。

	示例 :

	输入:

	   1
		\
		 3
		/
	   2

	输出:
	1

	解释:
	最小绝对差为1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
	注意: 树中至少有2个节点。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


var (
	prevalue, minvalue int
)
func getMinimumDifference(root *TreeNode) int {
	prevalue = -1
	minvalue = math.MaxInt64
	minOrder(root)
	return minvalue
}

func minOrder(root *TreeNode)  {
	if root == nil {
		return
	}
	minOrder(root.Left)
	if prevalue != -1 {
		if root.Val - prevalue < minvalue {
			minvalue = root.Val - prevalue
		}
	}
	prevalue = root.Val
	minOrder(root.Right)
}

func GetMinimumDifference(root *TreeNode) int {
	return getMinimumDifference(root)
}