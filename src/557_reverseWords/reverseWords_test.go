package _57_reverseWords

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	fmt.Println(ReverseWords("Let's take LeetCode contest"))
}


func TestReverseWords2(t *testing.T) {
	fmt.Println(ReverseWords("Let's      take LeetCode contest"))
}


func TestReverseWords3(t *testing.T) {
	fmt.Println(ReverseWords("    Let's      take LeetCode contest"))
}


func TestReverseWords4(t *testing.T) {
	fmt.Println(ReverseWords("    Let's      take LeetCode contest   "))
}

func TestReverseWords5(t *testing.T) {
	fmt.Println(ReverseWords(""))
}

func TestReverseWords6(t *testing.T) {
	fmt.Println(ReverseWords("               "))
}