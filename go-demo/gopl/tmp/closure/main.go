package main

import (
    "fmt"
)

func main() {
    var a int = 42
    c := make(chan struct{})
    go func () {
        for range c {
            fmt.Println(a)
        }
    }()
    c <- struct{}{}
    a = 43
    c <- struct{}{}

    close(c)

}
