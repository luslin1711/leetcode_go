package __StringToInteger

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	str := "  0000000000012345678"
	fmt.Println(MyAtoi(str))
}
