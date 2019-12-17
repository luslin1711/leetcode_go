package _01_isSymmetric

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if (root.Left == nil && root.Right == nil) {
		return true
	}
	if (root.Left != nil && root.Right == nil) || (root.Left == nil && root.Right != nil) {
		return false
	}
	return symmetric(root.Left,root.Right)
}

func symmetric(left, right *TreeNode) bool {
	var (
		left_res, right_res bool
	)
	if left.Val == right.Val{
		if left.Left != nil && right.Right != nil{
			left_res = symmetric(left.Left,right.Right)
		} else if (left.Left == nil && right.Right != nil) || (left.Left != nil && right.Right == nil)   {
			left_res= false
		} else {
			left_res= true
		}
		if left.Right != nil && right.Left != nil {
			right_res = symmetric(left.Right,right.Left)
		} else if (left.Right == nil && right.Left != nil) || (left.Right != nil && right.Left == nil) {
			right_res = false
		} else {
			right_res = true
		}
	} else {
		return false
	}
	return left_res && right_res
}

func IsSymmetric2(root *TreeNode) bool {
	return isSymmetric2(root)
}
