package _0_mergeTwoLists

import "fmt"

/*
	将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

	示例：

	输入：1->2->4, 1->3->4
	输出：1->1->2->3->4->4

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) Print() {
	for this != nil {
		fmt.Print(this.Val," ")
		this = this.Next
	}
	fmt.Println()
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		dummyHead *ListNode
	)
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummyHead = &ListNode{
		Val:  0,
		Next: l1,
	}
	l1 = dummyHead
	for l1.Next != nil && l2 != nil {
		if l1.Next.Val > l2.Val {
			l1.Next = &ListNode{Val:  l2.Val, Next: l1.Next,}
			l2 = l2.Next
		}
		l1 = l1.Next
	}
	for l1.Next != nil {
		l1 = l1.Next
	}
	for l2 != nil {
		l1.Next = &ListNode{
			Val:  l2.Val,
			Next: l2.Next,
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return dummyHead.Next
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeTwoLists(l1,l2)
}