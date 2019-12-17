package _8_lengthOfLastWord

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	fmt.Println(LengthOfLastWord("Hello World"))
}

func TestLengthOfLastWord1(t *testing.T) {
	fmt.Println(LengthOfLastWord("Hello World  "))
}

func TestLengthOfLastWord2(t *testing.T) {
	fmt.Println(LengthOfLastWord("Hello Wo rld  "))
}