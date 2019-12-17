package _01_isSymmetric

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(IsSymmetric(tree))
}

func TestIsSymmetric2(t *testing.T) {
	tree := &TreeNode{
		Val:   9,
		Left:  &TreeNode{
			Val:   -42,
			Left: nil,
			Right: &TreeNode{
				Val:   76,
				Left:  nil,
				Right: &TreeNode{
					Val:   13,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:  -42,
			Left: &TreeNode{
				Val:  76,
				Left: nil,
				Right: &TreeNode{
					Val:   13,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
	}
	fmt.Println(IsSymmetric2(tree))
}