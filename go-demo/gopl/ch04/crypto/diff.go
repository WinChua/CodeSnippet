package main

import (
    "fmt"
    "os"
    "crypto/sha256"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Printf("Usage %s str1 str2", os.Args[0])
        return
    }
    b32_1 := sha256.Sum256([]byte(os.Args[1]))
    b32_2 := sha256.Sum256([]byte(os.Args[2]))
    for i := range b32_1 {
        b32_1[i] ^= b32_2[i]
    }
    var s int
    for _, b := range b32_1 {
        var i uint
        for i = 0; i < 8; i++ {
            if b & (1 << i) == 1 {
                s++
            }
        }
    }

    fmt.Println(s)
}


