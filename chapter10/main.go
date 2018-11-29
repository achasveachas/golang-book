package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for index := 0; index < 10; index++ {
		fmt.Println(n, ":", index)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	// goroutines
	for index := 0; index < 10; index++ {
		go f(index)
	}
	var input string
	fmt.Scanln(&input)
}
