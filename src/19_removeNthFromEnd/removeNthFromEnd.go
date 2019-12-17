package _9_removeNthFromEnd

import "fmt"

/*
	给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
	示例：
	给定一个链表: 1->2->3->4->5, 和 n = 2.
	当删除了倒数第二个节点后，链表变为 1->2->3->5.
	说明：
	给定的 n 保证是有效的。
	进阶：
	你能尝试使用一趟扫描实现吗？
	在真实的面试中遇到过这道题？
	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func (this *ListNode) Print() {
	for this != nil {
		fmt.Println(this.Val)
		this = this.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		dummyHead *ListNode
		p,q *ListNode
	)
	if n == 0  || head.Next == nil{
		return head
	}
	dummyHead = &ListNode{
		Val:  122,
		Next: head,
	}
	p = dummyHead
	q = dummyHead
	for i:=0; i < n+1; i++ {
		p = p.Next
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return dummyHead.Next
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	return removeNthFromEnd(head,n)
}