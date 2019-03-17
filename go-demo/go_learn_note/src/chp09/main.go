package main

import (
    "fmt"
    "unsafe"
    "./visual"
    _ "./internal/ia"
    _ "./init"
)

func main() {
    d := visual.NewData(1, 2)
    fmt.Println(d)
    d.Y = 42
    fmt.Println(d)
    p := (*struct{x int})(unsafe.Pointer(d))
    p.x = 23
    fmt.Println(d)
}
