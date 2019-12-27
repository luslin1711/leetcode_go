package _80_validPalindrome

import (
	"fmt"
	"testing"
)


func TestValidPalindrome(t *testing.T) {
	fmt.Println(ValidPalindrome("aba"))
}

func TestValidPalindrome2(t *testing.T) {
	fmt.Println(ValidPalindrome("abca"))
}


func TestValidPalindrome3(t *testing.T) {
	fmt.Println(ValidPalindrome("adebbea"))
}


func TestValidPalindrome4(t *testing.T) {
	fmt.Println(ValidPalindrome("fefeffff"))
}

func TestValidPalindrome5(t *testing.T) {
	fmt.Println(ValidPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))
}