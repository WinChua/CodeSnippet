package main

import "fmt"

func no_return() (res int) {
    defer func() {
        recover()
        res = 42
    }()
    panic(42)
}

func main() {
    fmt.Println(no_return())
}

