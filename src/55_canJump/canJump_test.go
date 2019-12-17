package _5_canJump

import (
	"fmt"
	"testing"
)

func TestCanJump1(t *testing.T) {
	fmt.Println(CanJump([]int{2,3,1,1,4}))
}

func TestCanJump2(t *testing.T) {
	fmt.Println(CanJump([]int{3,2,1,0,4}))
}