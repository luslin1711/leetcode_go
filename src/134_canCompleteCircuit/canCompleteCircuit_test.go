package _34_canCompleteCircuit

import (
	"fmt"
	"testing"
)

func TestCanCompleteCircuit1(t *testing.T) {
	fmt.Println(CanCompleteCircuit([]int{1,2,3,4,5},[]int{3,4,5,1,2}))
}

func TestCanCompleteCircuit2(t *testing.T) {
	fmt.Println(CanCompleteCircuit([]int{2,3,4},[]int{3,4,3}))
}