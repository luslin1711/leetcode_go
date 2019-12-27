package _82_calPoints

import (
	"fmt"
	"testing"
)


func TestCalPoints(t *testing.T) {
	ops := []string{"5","2","C","D","+"}
	fmt.Println(CalPoints(ops))
}

func TestCalPoints2(t *testing.T) {
	ops := []string{"5","-2","4","C","D","9","+","+"}
	fmt.Println(CalPoints(ops))
}