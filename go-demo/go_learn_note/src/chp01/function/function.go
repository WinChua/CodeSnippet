package main

import (
    "fmt"
)

func add(x int) func(y int) int {
    return func(y int) int {
        return x + y
    }
}

func main() {
    addOne := add(1)
    addTwo := add(2)
    fmt.Println(addOne(3))
    fmt.Println(addTwo(2))
}
