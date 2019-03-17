package main

import (
    "fmt"
)

type inner struct {
    Age int
    Addr string
}

type outter struct {
    inner
    Name string
}

func (i inner) HelloWorld() {
    fmt.Printf("%T\n", i)
}

func main() {
    var o outter
    o.HelloWorld()
}
