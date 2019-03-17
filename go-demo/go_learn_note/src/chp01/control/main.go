package main

import (
    "strconv"
    "os"
    "fmt"
)

func main() {
    var num, _ = strconv.Atoi(os.Args[1])
    if (num < 42) {
        println("Less")
    } else {
        println("Hello")
    }

    switch {
    case num < 42:
        println("Switch less 42")
    case num < 42:
        println("oSwitch less 42")
    default:
        println("HelloSwitch")
    }

    for i := 1; i < 10; i++ {
        println(i)
    }

    list := make([]int, 10)
    println(list)
    list[0] = 42
    list = append(list, 1)
    for i, v := range list {
        println(i, v)
        list = append(list, i + v)
    }
    println(list)
    fmt.Println(list)

    x := 5
    for x > 0 {
        println(x)
        x--
    }
}
