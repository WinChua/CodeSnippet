package main

import (
    "fmt"
)

func main() {
    defer func(){
        switch p := recover(); p {
        case nil:
        case 1:
            fmt.Println(p)
        default:
            panic(p)
        }
    }()
    panic("hello")
}
