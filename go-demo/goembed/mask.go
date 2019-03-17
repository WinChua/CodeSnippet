package main

import "fmt"

type Interface interface {
    Read()
    Write()
}

type SubInterface interface {
    Read()
}

type II struct{}

func (i *II) Read() {
    fmt.Println("Read from II")
}
func (i *II) Write() {
    fmt.Println("Write from II")
}

type EE struct {
    SubInterface
    target * II
}

func NewEE() *EE {
    e := new(EE)
    e.target = new(II)
    e.SubInterface = e.target
    return e
}
