package _04_convertToBase7

import (
	"fmt"
	"testing"
)

func TestConvertToBase7(t *testing.T) {
	fmt.Println(ConvertToBase7(100))
}

func TestConvertToBase71(t *testing.T) {
	fmt.Println(ConvertToBase7(-100))
}


func TestConvertToBase72(t *testing.T) {
	fmt.Println(ConvertToBase7(-1))
}