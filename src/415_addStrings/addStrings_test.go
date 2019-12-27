package _15_addStrings

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T)  {
	fmt.Println("1"[0])
}

func TestAddStrings(t *testing.T) {
	fmt.Println(AddStrings("1234548","21561"))
}

func TestAddStrings1(t *testing.T) {
	fmt.Println(AddStrings("1234548","21561"))
}


func TestAddStrings2(t *testing.T) {
	fmt.Println(AddStrings("9999999","9999999"))
}