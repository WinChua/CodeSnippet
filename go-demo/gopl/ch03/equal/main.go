package main

import (
    "os"
    "fmt"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Printf("Usage %s str1 str2\n", os.Args[0])
        return
    }
    var str1 = os.Args[1]
    var str2 = os.Args[2]
    if str1 == str2 {
        fmt.Println("equal")
        return
    }
    count := make(map[rune]int)
    for _, r := range str1 {
        count[r]++
    }
    for _, r := range str2 {
        _, ok := count[r]
        if !ok {
            fmt.Println("No")
            return
        }
        count[r]--
    }
    for _, c := range count {
        if c != 0 {
            fmt.Println("No")
            return
        }
    }
    fmt.Println("Yes")

}
