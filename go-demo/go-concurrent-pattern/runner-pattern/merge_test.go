package main

import (
    "fmt"
    "testing"
    "./runner"
    "os"
)


func TestMerge(t *testing.T) {
    fmt.Println("the pid is", os.Getpid())
    c := runner.NewSignaler(os.Interrupt, os.Kill)
    for s := range c{
        fmt.Println(s)
    }
}

