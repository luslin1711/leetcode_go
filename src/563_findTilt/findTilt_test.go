package _63_findTilt

import (
	"fmt"
	"testing"
)

func TestFindTilt(t *testing.T) {
	tree := &TreeNode{
		Val:1,
		Left:&TreeNode{
			Val:2,
			Left:&TreeNode{
				Val:4,
			},
		},
		Right:&TreeNode{
			Val:3,
			Left:&TreeNode{
				Val:5,
			},
		},
	}
	fmt.Println(FindTilt(tree))
}