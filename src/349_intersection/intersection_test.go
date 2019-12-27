package _49_intersection

import (
	"fmt"
	"testing"
)

func TestIntersection(t *testing.T) {
	fmt.Println(Intersection([]int{1,2,2,1},[]int{2,2}))
}

func TestIntersection1(t *testing.T) {
	fmt.Println(Intersection([]int{4,9,5},[]int{9,4,9,8,4}))
}