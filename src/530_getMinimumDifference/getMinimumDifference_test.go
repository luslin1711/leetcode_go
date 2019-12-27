package _30_getMinimumDifference

import (
	"fmt"
	"testing"
)

func TestGetMinimumDifference(t *testing.T) {
	tree := &TreeNode{
		Val:1,
		Right:&TreeNode{
			Val:3,
			Left:&TreeNode{
				Val:2,
				Right:&TreeNode{
					Val:1,
				},
			},
		},
	}
	fmt.Println(GetMinimumDifference(tree))
}
