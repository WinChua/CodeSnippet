package main

import (
    "fmt"
    "os"
    "encoding/json"
    "sort"
)

type Person struct {
    Name string
    Age int
}

func main() {
    var data interface{}
    if len(os.Args) == 2 {
        data = []byte(`{"hello": "world"}`)
        fmt.Println(data)
    } else {
        data = Person{Name:"WinChua", Age: 26}
    }
    d, err := json.Marshal(data)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(d))
}
