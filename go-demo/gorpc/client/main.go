package main

import (
    "fmt"
    "../rpcargs"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Printf("Usage %s a b\n", os.Args[0])
        return
    }
    var a, b int
    a, _ = strconv.Atoi(os.Args[1])
    b, _ = strconv.Atoi(os.Args[2])
    c, err := rpcargs.NewClient()
    if err != nil {
        panic(err)
    }

    reply, err := c.Multiply(a, b)
    if err != nil {
        panic(err)
    }
    fmt.Println(reply)
}
