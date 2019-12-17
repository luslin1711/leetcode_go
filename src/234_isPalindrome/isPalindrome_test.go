package _34_isPalindrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	nodes := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	fmt.Println(IsPalindrome(nodes))
}

func TestIsPalindrome1(t *testing.T) {
	nodes := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(IsPalindrome(nodes))
}

func TestIsPalindrome2(t *testing.T) {
	nodes := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(IsPalindrome(nodes))
}