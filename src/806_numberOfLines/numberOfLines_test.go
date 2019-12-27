package _06_numberOfLines

import (
	"fmt"
	"testing"
)


func TestNumberOfLines(t *testing.T) {
	widths := []int{10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10}
	fmt.Println(numberOfLines(widths,"abcdefghijklmnopqrstuvwxyz"))
}

func TestNumberOfLines2(t *testing.T) {
	widths := []int{4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10}
	fmt.Println(numberOfLines(widths,"bbbcccdddaaa"))
}