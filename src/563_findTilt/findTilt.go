package _63_findTilt

import "math"

/*
	给定一个二叉树，计算整个树的坡度。

	一个树的节点的坡度定义即为，该节点左子树的结点之和和右子树结点之和的差的绝对值。空结点的的坡度是0。

	整个树的坡度就是其所有节点的坡度之和。

	示例:

	输入:
			 1
		   /   \
		  2     3
	输出: 1
	解释:
	结点的坡度 2 : 0
	结点的坡度 3 : 0
	结点的坡度 1 : |2-3| = 1
	树的坡度 : 0 + 0 + 1 = 1
	注意:

	任何子树的结点的和不会超过32位整数的范围。
	坡度的值不会超过32位整数的范围。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/binary-tree-tilt
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	d := 0
	RLV(root, &d)
	return d
}


func RLV(root *TreeNode, d *int) int {
	if root == nil {
		return 0
	}
	r := RLV(root.Right, d)
	l := RLV(root.Left, d)
	*d += int(math.Abs(float64(r-l)))
	return  root.Val + l + r
}

func FindTilt(root *TreeNode) int {
	return findTilt(root)
}