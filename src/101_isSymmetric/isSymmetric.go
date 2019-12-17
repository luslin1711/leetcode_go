package _01_isSymmetric

/*
	给定一个二叉树，检查它是否是镜像对称的。

	例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

		1
	   / \
	  2   2
	 / \ / \
	3  4 4  3
	但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

		1
	   / \
	  2   2
	   \   \
	   3    3
	说明:

	如果你可以运用递归和迭代两种方法解决这个问题，会很加分

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/symmetric-tree
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var (
		p_nodes, q_nodes []*TreeNode
		tmp_p, tmp_q *TreeNode
	)
	if root == nil {
		return false
	}
	if (root.Left == nil && root.Right == nil) {
		return true
	}
	if (root.Left != nil && root.Right == nil) || (root.Left == nil && root.Right != nil) {
		return false
	}
	p_nodes = make([]*TreeNode,0)
	q_nodes = make([]*TreeNode,0)
	p_nodes = append(p_nodes, root.Left)
	q_nodes = append(q_nodes,root.Right)
	for len(q_nodes) != 0 && len(p_nodes) != 0 {
		tmp_p = p_nodes[0]
		tmp_q = q_nodes[0]
		if (tmp_p.Val != tmp_q.Val)  {
			return false
		}
		if (tmp_p.Left == nil && tmp_q.Right != nil) || (tmp_p.Left != nil && tmp_q.Right == nil){
			return false
		} else if tmp_p.Left != nil && tmp_q.Right != nil {
			p_nodes = append(p_nodes,tmp_p.Left)
			q_nodes = append(q_nodes,tmp_q.Right)
		}

		if (tmp_p.Right == nil && tmp_q.Left != nil) || (tmp_p.Right != nil && tmp_q.Left == nil){
			return false
		} else if tmp_p.Right != nil && tmp_q.Left != nil {
			p_nodes = append(p_nodes,tmp_p.Right)
			q_nodes = append(q_nodes,tmp_q.Left)
		}
		p_nodes = p_nodes[1:]
		q_nodes = q_nodes[1:]
	}
	return true
}

func IsSymmetric(root *TreeNode) bool {
	return isSymmetric(root)
}