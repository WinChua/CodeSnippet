package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    str := make([]byte, 0, 100)
    str = strconv.AppendInt(str, 42, 10)
    str = strconv.AppendBool(str, true)
    str = strconv.AppendUint(str, 42, 2)
    str = strconv.AppendQuote(str, "quote")
    fmt.Println(strings.ToLower("HELLO, WORLD"))
    fmt.Println(string(str))
    b, err := strconv.ParseBool("false")
    if err != nil {
        panic(err)
    }
    fmt.Println(b)
}
