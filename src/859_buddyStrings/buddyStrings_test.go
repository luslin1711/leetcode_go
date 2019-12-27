package _59_buddyStrings

import (
	"fmt"
	"testing"
)


func TestBuddyStrings(t *testing.T) {
	a := "ab"
	b := "ba"
	fmt.Println(buddyStrings(a,b))
}


func TestBuddyStrings2(t *testing.T) {
	a := "ab"
	b := "ab"
	fmt.Println(buddyStrings(a,b))
}


func TestBuddyStrings3(t *testing.T) {
	a := "aa"
	b := "aa"
	fmt.Println(buddyStrings(a,b))
}


func TestBuddyStrings4(t *testing.T) {
	a := "ab"
	b := "ba"
	fmt.Println(buddyStrings(a,b))
}