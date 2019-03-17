package main

import (
    "os"
    "fmt"
    "strings"
)

func samestrip(data []string) []string{
    out := data[:1]
    for _, s := range data[1:] {
        if s == out[len(out)-1] {
            continue
        }
        out = append(out, s)
    }
    return out
}

func main() {
    data := strings.Split(os.Args[1], ",")
    data = samestrip(data)
    fmt.Println(data)
}


