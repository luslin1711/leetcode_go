package _57_binaryTreePaths

import (
	"fmt"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(BinaryTreePaths(tree))
}