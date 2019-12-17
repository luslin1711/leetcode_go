package _60_getIntersectionNode

import (
	"fmt"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	node := &ListNode{
		Val:  8,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	headA := &ListNode{
		Val:  4,
		Next: &ListNode{
			Val:  1,
			Next: node,
		},
	}
	headB := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  0,
			Next: &ListNode{
				Val:  1,
				Next: node,
			},
		},
	}
	fmt.Println(GetIntersectionNode(headA,headB).Val)
}
