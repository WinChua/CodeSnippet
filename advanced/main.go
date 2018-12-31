package main

import (
    "fmt"
    "time"
    "./wrap/fetch"
    "./wrap"
)


func main() {
    merged := wrap.Merge(
        wrap.Subscribe(fetch.Fetch("http://localhost:12340/getjson")),
        wrap.Subscribe(fetch.Fetch("http://localhost:12341/getjson")),
        wrap.Subscribe(fetch.Fetch("http://localhost:12341/getjson")),
    )
    time.AfterFunc(3 * time.Second, func() {
        fmt.Println("Closed:", merged.Close())
    })

    for it := range merged.Updates() {
        fmt.Println("The data is", it)
    }
}
