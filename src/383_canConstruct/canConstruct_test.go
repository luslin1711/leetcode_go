package _83_canConstruct

import (
	"fmt"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	fmt.Println(CanConstruct("a","b"))
}

func TestCanConstruct1(t *testing.T) {
	fmt.Println(CanConstruct("aa","ab"))
}


func TestCanConstruct2(t *testing.T) {
	fmt.Println(CanConstruct("aa","aab"))
}

func TestCanConstruct3(t *testing.T) {
	maps := map[int]int{}
	fmt.Println(len(maps))
}