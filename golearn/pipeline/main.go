package main

import (
    "fmt"
)

func main() {
    naturals := make(chan int)
    squarer := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            naturals <- i
        }
        close(naturals)
    }()

    go func() {
        for v := range naturals {
            squarer <- v*v
        }
        close(squarer)
    }()

    for v := range squarer {
        fmt.Println(v)
    }
}
