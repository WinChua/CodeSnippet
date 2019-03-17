package main

import (
    "os"
    "strings"
    "strconv"
    "fmt"
)

func reverse(s []int) {
    var i, j int
    for i, j = 0, len(s) - 1 ; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

func main() {
    var data []int
    for _, d := range strings.Split(os.Args[1], ",") {
        dd, _ := strconv.Atoi(d)
        data = append(data, dd)
    }
    fmt.Println(data)
    reverse(data)
    fmt.Println(data)
}
