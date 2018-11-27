# Problems

- How do you get the memory address of a variable?

For a variable `x` you get the memory address by referncing `&x`

- How do you assign a value to a pointer?

`*x = newValue`

- How do you create a new pointer?

Either with an asterisk followed by the type (e.g. `var x *int`) or using the `new` function (e.g. `x := new(int)`)

- What is the value of x after running this program:
```golang
func square(x *float64) {
  *x = *x * *x
}
func main() {
  x := 1.5
  square(&x)
}
```

`x` squared

- Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).

```golang
func swap(x *int, y *int) {
	tempX := *x
	*x = *y
	*y = tempX
}
```