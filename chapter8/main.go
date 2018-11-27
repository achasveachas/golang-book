package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	tempX := *x
	*x = *y
	*y = tempX
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0

	y := 2.0
	square(&y)
	fmt.Println(y)

	var1 := 1
	var2 := 2
	swap(&var1, &var2)
	fmt.Println(var1, var2)
}
