# Problems

- What does the following program print:

```golang
i := 10
if i > 10 {
  fmt.Println("Big")
} else {
  fmt.Println("Small")
}
```

`Small`

- Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)

The way they expect you to write it:

```golang
for i := 1; i < 100; i++ {
  if i%3 == 0 {
    fmt.Println(i)
  }
}
```

[Badass](https://blog.yechiel.me/for-loops-in-javascript-cda0388b6dd0) way to write it ;)

```golang
for i := 3; i < 10; i += 3 {
  if i%3 == 0 {
    fmt.Println(i)
  }
}
```

- Write a program that prints the numbers from 1 to 100. But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz". For numbers which are multiples of both three and five print "FizzBuzz".

```golang
for i := 0; i < 20; i++ {
  if i%3 == 0 && i%5 == 0 {
    fmt.Println("Fizzbuzz")
  } else if i%3 == 0 {
    fmt.Println("Fizz")
  } else if i%5 == 0 {
    fmt.Println("Buzz")
  } else {
    fmt.Println(i)
  }
}
```