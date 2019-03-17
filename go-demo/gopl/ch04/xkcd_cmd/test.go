package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup

func producer(name string, dataCh chan<- string) {
    defer wg.Done()
    dataCh <- name + "Hello"
}

func consumer(name string, dataCh <-chan string) {
    for value := range dataCh {
        fmt.Println(name, "get", value)
    }
}

func main() {
    dataCh := make(chan string, 2)
    for i:=0; i<10; i++ {
        wg.Add(1)
        go producer(fmt.Sprintf("P%d", i), dataCh)
    }

    go func() {
        wg.Wait()
        close(dataCh)
    }()

    consumer("C0", dataCh)
    fmt.Println("End")
}
