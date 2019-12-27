package _66_matrixReshape

import (
	"fmt"
	"testing"
)

func TestMatrixReshape(t *testing.T) {
	nums := [][]int{{1,2},{3,4},{5,6}}
	fmt.Println(MatrixReshape(nums,6,1))
}