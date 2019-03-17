package main

import (
    "flag"
    "crypto/sha256"
    "crypto/sha512"
    "fmt"
)

var which string

func init() {
    flag.StringVar(&which, "which", "256", "choice which sha to use")
}

func main() {
    flag.Parse()
    data := flag.Arg(0)
    switch which {
    case "256":
        fmt.Printf("%s: %x\n", data, sha256.Sum256([]byte(data)))
    case "512":
        fmt.Printf("%s: %x\n", data, sha512.Sum512([]byte(data)))
    case "384":
        fmt.Printf("%s: %x\n", data, sha512.Sum384([]byte(data)))
    default:
        fmt.Println("Error: -which should in 256, 512 or 384")
    }
}
