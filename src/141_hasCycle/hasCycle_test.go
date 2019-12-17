package _41_hasCycle

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	nodelast := &ListNode{
		Val:  5,
		Next: nil,
	}
	nodes := &ListNode{
		Val:  3,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  0,
				Next: &ListNode{
					Val:  -4,
					Next: nodelast,
				},
			},
		},
	}
	nodelast.Next = nodes
	fmt.Println(HasCycle(nodes))


}