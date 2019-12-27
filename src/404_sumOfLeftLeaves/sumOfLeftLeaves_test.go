package _04_sumOfLeftLeaves

import (
	"fmt"
	"testing"
)

func TestSumOfLeftLeaves(t *testing.T) {
	tree := &TreeNode{
		Val:3,
		Left:&TreeNode{Val:9},
		Right:&TreeNode{Val:20,Left:&TreeNode{Val:15},Right:&TreeNode{Val:7}},
	}
	fmt.Println(SumOfLeftLeaves(tree))
}