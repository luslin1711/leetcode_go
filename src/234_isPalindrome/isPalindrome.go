package _34_isPalindrome


/*
	请判断一个链表是否为回文链表。

	示例 1:

	输入: 1->2
	输出: false
	示例 2:

	输入: 1->2->2->1
	输出: true
	进阶：
	你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/palindrome-linked-list
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type ListNode struct {
	Val int
	Next *ListNode
}


func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slowp, fastp := head, head
	for fastp != nil && fastp.Next != nil {
		fastp = fastp.Next.Next
		slowp = slowp.Next
	}
	// 从slowp 开始反转链表
	preNode := slowp
	tmp := slowp.Next
	for tmp != nil {
		nextNode := tmp.Next
		tmp.Next = preNode
		preNode = tmp
		tmp = nextNode
	}
	for head != slowp {
		if head.Val != preNode.Val {
			return false
		}
		head = head.Next
		preNode = preNode.Next
	}
	return true
}


func IsPalindrome(head *ListNode) bool {
	return isPalindrome(head)
}