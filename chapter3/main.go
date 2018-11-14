package main

import "fmt"

func main() {
	fmt.Println("Int operations:")
	fmt.Println("1 + 1 =", 1 + 1)

	fmt.Println(len("String Operations"))
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	fmt.Println(len("Bool Operations"))
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}