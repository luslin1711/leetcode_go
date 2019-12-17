package _0_mergeTwoLists

import "testing"

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		Val:  -9,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	}
	l2 := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  7,
			Next: nil,
		},
	}
	mergeTwoLists(l1,l2).Print()
}
