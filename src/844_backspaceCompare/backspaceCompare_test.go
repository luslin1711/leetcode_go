package _44_backspaceCompare

import (
	"fmt"
	"testing"
)


func TestBackspaceCompare(t *testing.T) {
	s := "ab#c"
	tq := "ad#c"
	fmt.Println(backspaceCompare(s,tq))
}

func TestBackspaceCompare1(t *testing.T) {
	s := "ab##"
	tq := "c#d#"
	fmt.Println(backspaceCompare(s,tq))
}

func TestBackspaceCompare2(t *testing.T) {
	s := "a##c"
	tq := "#a#c"
	fmt.Println(backspaceCompare(s,tq))
}





func TestBackspaceCompare3(t *testing.T) {
	s := "bxj##tw"
	tq := "bxo#j##tw"
	fmt.Println(backspaceCompare(s,tq))
}