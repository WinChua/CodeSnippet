package main

import "fmt"

func counter(out chan<- int) {
    for i:=0; i<100; i++ {
        out <- i
    }
    close(out)
}

func squarer(out chan<- int, in <-chan int) {
    for v := range in {
        out <- v * v
    }
    close(out)
}

func printer(in  <-chan int) {
    for v := range in {
        fmt.Printf("%d ", v)
    }
    fmt.Println("")
}

func main() {
    n := make(chan int)
    s := make(chan int)
    go counter(n)
    go squarer(s, n)
    printer(s)
}
