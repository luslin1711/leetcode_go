package _03_removeElements

import "fmt"

/*
	删除链表中等于给定值 val 的所有节点。

	示例:

	输入: 1->2->6->3->4->5->6, val = 6
	输出: 1->2->3->4->5
 */


type ListNode struct {
	Val int
	Next *ListNode
}

func (this *ListNode) Print(){
	for this != nil {
		fmt.Print(this.Val," ")
		this = this.Next
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyhead := &ListNode{
		Val:  -1<<31,
		Next: head,
	}
	originHead := dummyhead
	preHead := dummyhead
	dummyhead = dummyhead.Next
	for dummyhead!= nil {
		if dummyhead.Val == val {
			preHead.Next = dummyhead.Next
		} else {
			preHead = preHead.Next
		}

		dummyhead = dummyhead.Next
	}
	return originHead.Next
}

func RemoveElements(head *ListNode, val int) *ListNode {
	return removeElements(head,val)
}