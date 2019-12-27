package _38_convertBST

import "testing"

func TestConvertBST(t *testing.T) {
	tree := &TreeNode{
		Val:5,
		Left:&TreeNode{
			Val:2,
		},
		Right:&TreeNode{
			Val:13,
		},
	}
	ConvertBST(tree)
}
