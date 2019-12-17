package _3_deleteDuplicates

import "fmt"

/*
	给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

	示例 1:

	输入: 1->1->2
	输出: 1->2
	示例 2:

	输入: 1->1->2->3->3
	输出: 1->2->3

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func (this *ListNode) Print() {
	for this != nil {
		fmt.Print(this.Val, " ")
		this = this.Next
	}
	fmt.Println()
}
func deleteDuplicates(head *ListNode) *ListNode {
	var (
		originHead *ListNode
		nextNode *ListNode
		nextNextNode *ListNode
	)
	if head == nil {
		return nil
	}
	originHead = head
	nextNode = head.Next
	for nextNode != nil {
		if nextNode.Val == head.Val {
			nextNextNode = nextNode.Next
			nextNode.Next = nil
			head.Next = nextNextNode
			nextNode = nextNextNode
		} else {
			head = head.Next
			nextNode = nextNode.Next
		}

	}
	return originHead
}

func DeleteDuplicates(head *ListNode) *ListNode {
	return deleteDuplicates(head)
}