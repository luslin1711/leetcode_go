package _90_wordPattern

import (
	"fmt"
	"testing"
)

func TestWordPattern(t *testing.T) {
	fmt.Println(WordPattern("abba","dog cat cat dog"))
}

func TestWordPattern1(t *testing.T) {
	fmt.Println(WordPattern("abba","dog cat cat fish"))
}

func TestWordPattern2(t *testing.T) {
	fmt.Println(WordPattern("aaaa","dog cat cat dog"))
}

func TestWordPattern3(t *testing.T) {
	fmt.Println(WordPattern("abba", "dog dog dog dog"))
}