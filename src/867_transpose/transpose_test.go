package _67_transpose

import (
	"fmt"
	"testing"
)


func TestTranspose(t *testing.T) {
	fmt.Println(transpose([][]int{{1,2,3},{4,5,6},{7,8,9}}))
}

func TestTranspose2(t *testing.T) {
	fmt.Println(transpose([][]int{{1,2,3},{4,5,6}}))
}