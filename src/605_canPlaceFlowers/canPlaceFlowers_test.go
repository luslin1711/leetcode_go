package _05_canPlaceFlowers

import (
	"fmt"
	"testing"
)


func TestCanPlaceFlowers(t *testing.T) {
	fmt.Println(CanPlaceFlowers([]int{1,0,0,0,1},1))
}

func TestCanPlaceFlowers2(t *testing.T) {
	fmt.Println(CanPlaceFlowers([]int{0,0,1,0,1},1))
}

func TestCanPlaceFlowers3(t *testing.T) {
	fmt.Println(CanPlaceFlowers([]int{0,0,1,0,0},3))
}