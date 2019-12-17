package _3_deleteDuplicates

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	head = DeleteDuplicates(head)
	head.Print()
}