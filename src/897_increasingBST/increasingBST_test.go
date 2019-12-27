package _97_increasingBST

import "testing"


func TestIncreasingBST(t *testing.T) {
	tree := &TreeNode{
		Val:5,
		Left:&TreeNode{
			Val:3,
			Left:&TreeNode{Val:2},
		},
		Right:&TreeNode{
			Val:6,
			Right:&TreeNode{
				Val:8,
			},
		},
	}
	increasingBST(tree)
}