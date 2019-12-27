package _32_flipAndInvertImage

import (
	"fmt"
	"testing"
)


func TestFlipAndInvertImage(t *testing.T) {
	image := [][]int{{1,0,0},{0,1,0},{1,1,1}}
	fmt.Println(flipAndInvertImage(image))
}

func TestFlipAndInvertImage2(t *testing.T) {
	image := [][]int{{1,1,0,0},{1,0,0,1},{0,1,1,1},{1,0,1,0}}
	fmt.Println(flipAndInvertImage(image))
}