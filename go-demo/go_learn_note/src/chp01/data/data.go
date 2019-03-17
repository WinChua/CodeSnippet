package main

import (
    "fmt"
)

func mapshow() {
    x := make(map[int]string)
    println(x, len(x))
    x[1] = "string"
    println(x, len(x))
    fmt.Println(x)
    delete(x, 1)
    fmt.Println(x)
    x[1] = "HelloWorld"
    fmt.Println(x[1])
    d, ok := x[1]
    fmt.Println(d, ok)
    d = x[1]
    d = x[2]
    fmt.Println(d)
    d, ok =  x[2]
    fmt.Println(d, ok)
}

func slice() {
    x := make([]int, 0, 1)
    println(x)
    println(x, len(x), cap(x))
    for i := 0; i < 40; i++ {
        x = append(x, i)
        println(i, x, len(x), cap(x))
    }
}

func main() {
    fmt.Println("slice()")
    slice()
    fmt.Println("mapshow()")
    mapshow()
}
