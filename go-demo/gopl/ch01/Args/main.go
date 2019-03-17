package main

import (
    "os"
    "fmt"
)

func main() {
    fmt.Println("The os.Args is ", os.Args)
    fmt.Println("The os.Args is a slice of string.")
    fmt.Println("The len, cap of os.Args is ", len(os.Args), cap(os.Args))
    for i, s := range os.Args {
        fmt.Println(i, s)
    }

}
