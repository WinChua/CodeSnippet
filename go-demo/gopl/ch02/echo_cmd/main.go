package main

import (
    "fmt"
    "strings"
    "flag"
)

var n bool
var sep string

func init() {
    flag.BoolVar(&n, "n", false, "omit the newline")
    flag.StringVar(&sep, "sep", " ", "specify the sep char")
}

func main() {
    flag.Parse()
    data := strings.Join(flag.Args(), sep)
    if n {
        fmt.Print(data)
    } else {
        fmt.Println(data)
    }
}


