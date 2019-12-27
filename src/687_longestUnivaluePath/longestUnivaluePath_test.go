package _87_longestUnivaluePath

import (
	"fmt"
	"testing"
)


func TestLongestUnivaluePath(t *testing.T) {
	tree := &TreeNode{
		Val:5,
		Left:&TreeNode{
			Val:4,
			Left:&TreeNode{
				Val:4,
			},
			Right:&TreeNode{
				Val:4,
			},
		},
		Right:&TreeNode{
			Val:5,
			Right:&TreeNode{Val:5},
		},
	}
	fmt.Println(longestUnivaluePath(tree))
}

