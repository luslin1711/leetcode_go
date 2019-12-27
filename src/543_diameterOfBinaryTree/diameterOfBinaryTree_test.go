package _43_diameterOfBinaryTree

import (
	"fmt"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	tree := &TreeNode{
		Val:1,
		Left:&TreeNode{
			Val:2,
			Left:&TreeNode{
				Val:4,
			},
			Right:&TreeNode{
				Val:5,
			},
		},
		Right:&TreeNode{Val:3},
	}
	fmt.Println(DiameterOfBinaryTree(tree))
}

func TestDiameterOfBinaryTree2(t *testing.T) {
	tree := &TreeNode{
		Val:1,
		Left:&TreeNode{
			Val:2,}}
	fmt.Println(DiameterOfBinaryTree(tree))
}