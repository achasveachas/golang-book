package main

import (
	"fmt"
)

func main() {
	// Arrays
	x := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	// Slices
	arr := [5]float64{1, 2, 3, 4, 5}
	y := arr[:]

	fmt.Println(y)

	// Slice functions
	y = append(y, 6, 7)
	fmt.Println(y)

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	// Maps
}
