package _221_balancedStringSplit

import (
	"fmt"
	"testing"
)

func TestBalancedStringSplit1(t *testing.T) {
	fmt.Println(BalancedStringSplit("RLRRLLRLRL"))
}

func TestBalancedStringSplit2(t *testing.T) {
	fmt.Println(BalancedStringSplit("RLLLLRRRLR"))
}

func TestBalancedStringSplit3(t *testing.T) {
	fmt.Println(BalancedStringSplit("LLLLRRRR"))
}