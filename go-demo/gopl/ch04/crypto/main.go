package main

import (
    "fmt"
    "os"
    "crypto/sha256"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Printf("Usage %s s...", os.Args[0])
        return
    }

    for _, s := range os.Args[1:] {
        r := sha256.Sum256([]byte(s))
        fmt.Printf("%s: %x\n", s, r)
    }

}


