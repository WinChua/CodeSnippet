package main

import (
    "fmt"
)

func getKBMB(i uint) uint64 {
    var s uint64 = 1
    for ; i > 0; i-- {
        s *= 1000
    }
    return s
}
func main() {
    fmt.Println(getKBMB(1), getKBMB(2))
}
