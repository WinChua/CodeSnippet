package main

import (
    "fmt"
)

func main() {
    producer1 := make(chan struct{})
    producer2 := make(chan struct{})
    consumer := make(chan int)
    go func() {
        defer close(consumer)
        defer close(producer1)
        for i:=0; i<=100; i+=2{
            <-producer2
            consumer <- i
            producer1 <- struct{}{}
        }
    }()
    go func() {
        defer close(producer2)
        for i:=1; i<=99; i+=2{
            <-producer1
            consumer <- i
            producer2 <- struct{}{}
        }
        <-producer1
    }()
    producer2 <- struct{}{}
    for i := range consumer {
        fmt.Println(i)
    }
}
