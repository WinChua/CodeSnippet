package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    abort := make(chan struct{})
    go func() {
        os.Stdin.Read(make([]byte,1))
        abort <- struct{}{}
    }()

    fmt.Println("lanch after 10s. or return to abort")
    select {
    case <-time.After(10 * time.Second):
        fmt.Println("lanch")
    case <-abort:
        fmt.Println("abort")
        return
    }
}
