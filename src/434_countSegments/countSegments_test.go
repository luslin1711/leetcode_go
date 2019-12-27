package _34_countSegments

import (
	"fmt"
	"testing"
)

func TestCountSegments(t *testing.T) {
	fmt.Println(CountSegments( "Hello, my name is John"))
}


func TestCountSegments2(t *testing.T) {
	fmt.Println(CountSegments( "  Hello, my name is John"))
}


func TestCountSegments3(t *testing.T) {
	fmt.Println(CountSegments( "  Hello, my name is John  "))
}

func TestCountSegments4(t *testing.T) {
	fmt.Println(CountSegments( "  Hello,   m y  name is John  "))
}