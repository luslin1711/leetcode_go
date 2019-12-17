package _00_isSameTree

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	p := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	q := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(IsSameTree(p,q))
}
