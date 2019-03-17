package main

import (
    "fmt"
)

type Data int64

func (x Data) hello(d int) {
    fmt.Println("Hello from Data")
    fmt.Println(d)
}

type He interface {
    hello(int)
}

func main() {
    var h He;
    var d Data = 10
    d.hello(10)
    h = d
    h.hello(42)
}
