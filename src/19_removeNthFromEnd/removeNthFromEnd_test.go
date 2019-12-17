package _9_removeNthFromEnd

import "testing"

func TestRemoveNthFromEnd1(t *testing.T) {
	lists := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},}

	RemoveNthFromEnd(&lists,2)
	lists.Print()
}

func TestRemoveNthFromEnd2(t *testing.T) {
	lists := ListNode{
		Val:  1,
		Next: nil}

	RemoveNthFromEnd(&lists,1)
	lists.Print()
}