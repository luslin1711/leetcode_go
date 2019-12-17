package __AddTwoNumbers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) Print() {
	for this != nil {
		fmt.Printf("%d->",this.Val)
		this = this.Next
	}
	fmt.Println("end")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		curr,dummyHead *ListNode
		carry int
		x, y, sum int
	)
	carry = 0
	dummyHead = &ListNode{
		Val:  0,
		Next: nil,
	}
	curr = dummyHead
	for l1 != nil || l2 != nil {
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		} else {
			x = 0
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		} else {
			y = 0
		}
		sum = carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		curr = curr.Next

	}
	if carry > 0 {
		curr.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return dummyHead.Next
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode  {
	return addTwoNumbers(l1,l2)
}
