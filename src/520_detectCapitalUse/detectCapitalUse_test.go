package _20_detectCapitalUse

import (
	"fmt"
	"testing"
)

func TestDetectCapitalUse(t *testing.T) {
	fmt.Println(DetectCapitalUse("FlaG"))
}


func TestDetectCapitalUse1(t *testing.T) {
	fmt.Println(DetectCapitalUse("USA"))
}

func TestDetectCapitalUse2(t *testing.T) {
	fmt.Println(DetectCapitalUse("aaaaaa"))
}
