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

func TestBalancedStringSplit(t *testing.T) {
	ans := make([]int,0)
	for i := 1; i < 10001; i++ {
		tmp := i
		for tmp != 0 {
			bit := tmp % 10
			if bit == 0 || tmp % bit != 0 {
				goto end
			}
			tmp /= 10
		}
		ans = append(ans, i)
	end:
	}
	fmt.Println(ans)
}