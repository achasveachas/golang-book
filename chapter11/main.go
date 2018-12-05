package main

import (
	"fmt"
	"golang-book/chapter11/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Avarage(xs)
	fmt.Println(avg)
}
