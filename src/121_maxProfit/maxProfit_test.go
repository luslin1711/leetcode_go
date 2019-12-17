package _21_maxProfit

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	fmt.Println(MaxProfit([]int{7,1,5,3,6,4}))
}


func TestMaxProfit1(t *testing.T) {
	fmt.Println(MaxProfit([]int{7,6,5,4,3,2}))
}

func TestMaxProfit2(t *testing.T) {
	fmt.Println(MaxProfit([]int{7}))
}

func TestMaxProfit3(t *testing.T) {
	fmt.Println(MaxProfit([]int{}))
}
