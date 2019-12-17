package __ZigZagConversion

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	fmt.Println(Convert("PAYPALISHIRING",3))
	fmt.Println(Convert("PAYPALISHIRING",4))
	fmt.Println(Convert("PAYPALISHIRING",100))
}
