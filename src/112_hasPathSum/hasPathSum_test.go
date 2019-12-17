package _12_hasPathSum

import (
	"fmt"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	tree := &TreeNode{
		Val:   5,
		Left:  &TreeNode{
			Val:   4,
			Left:  &TreeNode{
				Val:   11,
				Left:  &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	fmt.Println(HasPathSum(tree,21))
}