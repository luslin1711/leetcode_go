package _50_intersect

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	fmt.Println(Intersect([]int{1,2,2,1},[]int{2,2}))
}


func TestIntersect2(t *testing.T) {
	fmt.Println(Intersect([]int{4,9,5},[]int{9,4,9,8,4}))
}