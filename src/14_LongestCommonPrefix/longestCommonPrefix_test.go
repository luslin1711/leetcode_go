package _4_LongestCommonPrefix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix1(t *testing.T) {
	strs := []string{"flower","flow","flight"}
	fmt.Println(LongestCommonPrefix(strs))
}

func TestLongestCommonPrefix2(t *testing.T) {
	strs := []string{"dog","racecar","car"}
	fmt.Println(LongestCommonPrefix(strs))
}