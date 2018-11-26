# Problems

- `sum` is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?

```golang
func sum(ints ...int) (int)

```

- Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example `half(1)` should return `(0, false)` and `half(2)` should return `(1, true)`.

```golang
	half := func(n int) (int, bool) {
		half := int(n / 2)
		isEven := n%2 == 0

		return half, isEven
	}
```

- Write a function with one variadic parameter that finds the greatest number in a list of numbers.

```golang
func greatest (list ...int) int {
    greatest := 0
    for _, v := range list {
        if v > greatest {
            greatest = v
        }
    }
    return greatest
}
```
- Using `makeEvenGenerator` as an example, write a `makeOddGenerator` function that generates odd numbers.

```golang
func makeOddGenerator() func() uint {
  i := uint(1)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}
```

- The Fibonacci sequence is defined as: `fib(0) = 0`, `fib(1) = 1`, `fib(n) = fib(n-1) + fib(n-2)`. Write a recursive function which can find `fib(n)`.

```golang
func fib(x uint) uint {
  if x == 0 || x == 1 {
    return x
  } else {
      return fib(x-1) + fib(x-2)
  }
}
```

- What are `defer`, `panic` and `recover`? How do you recover from a run-time panic?

`defer` runs at the end of the function, `panic` raises a runtime error, `recover` rescues the `panic`.