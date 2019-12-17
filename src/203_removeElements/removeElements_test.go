package _03_removeElements

import (
	"testing"
)

func TestRemoveElements(t *testing.T) {
	node := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  6,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
							Val:  5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	RemoveElements(node,6).Print()
}