package main

import (
    "fmt"
    "encoding/json"
)

type G struct {
    D map[int]string
}

func (g G) Change(d string) {
    g.D[1] = d
}

type H struct {
    Hello string `json:"hello"`
}

func main() {
    fmt.Println("Hello, World")
    data := map[int]string {1: "hello"}
    g := G{D:data}
    fmt.Println(data, g)
    g.Change("World")
    fmt.Println(data, g)
    in := `{"hello": "world"}`
    //rawIn := json.RawMessage(in)
    bytes, err := json.Marshal([]byte(in))
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(bytes))

    h := new(H)
    err = json.Unmarshal([]byte(in), h)
    if err != nil {
        panic(err)
    }
    fmt.Println(h)
}
