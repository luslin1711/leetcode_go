package __AddTwoNumbers

import "testing"

func TestAddTwoNumbers1(t *testing.T) {
	l1 := &ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	l2 := &ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	res := AddTwoNumbers(l1,l2)
	res.Print()
}

func TestAddTwoNumbers2(t *testing.T) {
	l1 := &ListNode{}
	l2 := &ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	res := AddTwoNumbers(l1,l2)
	res.Print()
}

func TestAddTwoNumbers3(t *testing.T) {
	l1 := &ListNode{
		Val:  9,
		Next: &ListNode{
			Val:  9,
			Next: nil,
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: nil,
	}
	res := AddTwoNumbers(l1,l2)
	res.Print()
}