package _99_findRestaurant

import (
	"fmt"
	"testing"
)


func TestFindRestaurant(t *testing.T) {
	fmt.Println(FindRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
								[]string{"KFC", "Shogun", "Burger King"}))
}


func TestFindRestaurant2(t *testing.T) {
	fmt.Println(FindRestaurant([]string{"Shogun", "KFC", "Burger King"},
		[]string{"KFC", "Shogun", "Burger King"}))
}