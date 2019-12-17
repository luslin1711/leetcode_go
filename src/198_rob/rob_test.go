package _98_rob

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	fmt.Println(Rob([]int{1,2,3,1}))
}

func TestRob1(t *testing.T) {
	fmt.Println(Rob([]int{2,7,9,3,1}))
}

func TestRob2(t *testing.T) {
	fmt.Println(Rob([]int{1}))
}

func TestRob3(t *testing.T) {
	fmt.Println(Rob([]int{1,2,0,0,0,0,7,1,4,1}))
}