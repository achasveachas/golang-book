Problems
- What are two ways to create a new variable?

`var x string = "Hello"` and `x := "Hello"`

- What is the value of x after running:
`x := 5; x += 1`?

`6`

- What is scope and how do you determine the scope of a variable in Go?

Scope is where in the code you have access to a given variable. In Go scope is determined by blocks, i'e you have access to a variable anywhere within the curly brackets it was declared in.

- What is the difference between var and const?

Var can be reassigned, const can't

- Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)

```golang
func main(f int) {
  c := (f - 32) * (5/9)
  return c
}

```
- Write another program that converts from feet into meters. (1 ft = 0.3048 m)

```golang
func main(feet int) {
  meters := feet * 0.3048
  return c
}

```