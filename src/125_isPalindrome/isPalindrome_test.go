package _25_isPalindrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
}

func TestIsPalindrome2(t *testing.T) {
	fmt.Println(IsPalindrome("0P"))
}