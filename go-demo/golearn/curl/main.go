package main

import (
    "fmt"
    "os"
    "net/http"
    "io"
)

func main() {
    resp , err := http.Get(os.Args[1])
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }
    file, err := os.Create(os.Args[2])
    if err != nil {
        fmt.Println("ERROR:", err)
        return
    }
    defer file.Close()

    dest := io.MultiWriter(file, os.Stdout)
    io.Copy(dest, resp.Body)
    if err := resp.Body.Close(); err != nil {
        fmt.Println("ERROR", err)
    }
}
