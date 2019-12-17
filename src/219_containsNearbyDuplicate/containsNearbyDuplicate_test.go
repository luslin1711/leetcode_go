package _19_containsNearbyDuplicate

import (
	"fmt"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	fmt.Println(ContainsNearbyDuplicate([]int{1,2,3,1},3))
}

func TestContainsNearbyDuplicate1(t *testing.T) {
	fmt.Println(ContainsNearbyDuplicate([]int{1,2,3,1},2))
}