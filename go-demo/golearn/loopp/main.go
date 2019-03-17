package main

import "fmt"

func main() {
    slice := make([]int, 0, 10)
    for i:=0; i<10; i++ {
        slice = append(slice, i*i)
    }

    ch := make(chan struct{})
    for _, v := range slice {
        go func(v int) {
            fmt.Println(v)
            ch <- struct{}{}
        }(v)
    }

    for range slice {
        <-ch
    }

}
