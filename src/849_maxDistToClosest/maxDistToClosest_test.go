package _49_maxDistToClosest

import (
	"fmt"
	"testing"
)


func TestMaxDistToClosest(t *testing.T) {
	fmt.Println(maxDistToClosest([]int{1,0,0,0,1,0,1}))
}

func TestMaxDistToClosest2(t *testing.T) {
	fmt.Println(maxDistToClosest([]int{1,0,0,0}))
}

func TestMaxDistToClosest3(t *testing.T) {
	fmt.Println(maxDistToClosest([]int{1,0,0,0,1,1,0,0,0,0,0,1,0,0,0,0}))
}

func TestMaxDistToClosest4(t *testing.T) {
	fmt.Println(maxDistToClosest([]int{1,0,1}))
}

func TestMaxDistToClosest5(t *testing.T) {
	fmt.Println(maxDistToClosest([]int{0,1,1,0,1,1,0,1}))
}
