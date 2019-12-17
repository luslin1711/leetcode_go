package _06_reverseList

import "testing"

func TestReverseList(t *testing.T) {
	nodes := &ListNode{
		Val:  3,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  0,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}
	ReverseList(nodes).Print()
}

func TestReverseList2(t *testing.T) {
	nodes := &ListNode{
		Val:  3,
		Next: nil,
	}
	ReverseList(nodes).Print()
}