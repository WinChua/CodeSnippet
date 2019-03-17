package main

import (
    "os"
    "fmt"
    "io/ioutil"
)

func MkdirAndRmdir() {
    err := os.Mkdir("hello", 0640)
    if err != nil {
        fmt.Println("Mkdir Error:", err)
        return
    }
    defer os.Remove("hello")

    err = os.MkdirAll("world/hello/winchua", 0777)
    if err != nil {
        fmt.Println("MkdirAll Error:", err)
        return
    }
    if tmp, err := ioutil.ReadDir("."); err == nil {
        for _, f := range tmp {
            fmt.Println(f.Name())
        }
    }

    defer os.RemoveAll("world")
}
