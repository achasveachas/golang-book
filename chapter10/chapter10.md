# Problems

- How do you specify the direction of a channel type?

`<- channel` for outgoing, `channel <-` for incoming.

- Write your own Sleep function using time.After.

```golang
func Sleep(timeOut int) {
    time.After(time.Second * time.Duration(timeOut))
}
```

- What is a buffered channel? How would you create one with a capacity of 20?

A channel with a specific capacity.

`make(chan int, 20)`