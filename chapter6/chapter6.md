# Problems

-How do you access the 4th element of an array or slice?
`arr[3]`

-What is the length of a slice created using: make([]int, 3, 9)?
3

-Given the array:
`x := [6]string{"a","b","c","d","e","f"}`
what would `x[2:5]` give you?
A slice that looks like `["c","d","e","f"]`

-Write a program that finds the smallest number in this list:

```golang
x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}
```

```golang
min := 0

for _, num := range x {
    if min == 0 || min > num {
        min = num
    }
}

return num
```