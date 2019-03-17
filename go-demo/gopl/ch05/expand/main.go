package main

import (
    "fmt"
    "strings"
    "os"
    "log"
)

func expand(str, pattern string, f func(string)string) string {
    newStr := f(pattern)
    return strings.Replace(str, pattern, newStr, -1)
}

func main() {
    if len(os.Args) != 3 {
        log.Fatalf("Usage %s string pattern", os.Args[0])
        return
    }

    str := os.Args[1]
    pattern := os.Args[2]

    fmt.Println(expand(str, pattern, strings.ToUpper))
    fmt.Println(expand(str, pattern, strings.Title))

}
