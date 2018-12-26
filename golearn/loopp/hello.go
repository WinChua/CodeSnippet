package main

import (
    "time"
    "fmt"
    "os"
)

func main() {
    abort := make(chan struct{})
    go func() {
        os.Stdin.Read(make([]byte, 1))
        abort <- struct{}{}
    }()
    fmt.Println("Start...")
    tick := time.Tick(1 * time.Second)
    for countDown := 10; countDown > 0; countDown-- {
        fmt.Println(<-tick)
    }
}


