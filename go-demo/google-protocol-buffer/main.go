package main

import (
    "fmt"
    "./tutorial"
    "github.com/golang/protobuf/proto"
    "encoding/json"
)

func main() {
    p := tutorial.Person {
        Id: 1234,
        Name: "WinChua",
        Email: "winchua@foxmail.com",
        Phones: []*tutorial.Person_PhoneNumber {
            {Number: "13660270454", Type: tutorial.Person_MOBILE},
        },
    }
    a := tutorial.Test1 {A: 150}
    fmt.Println(p)
    out, err := proto.Marshal(&p)
    if err != nil {
        fmt.Println(err)
        return
    }
    aout, _ := proto.Marshal(&a)
    fmt.Println("the len of pb test1:", len(aout))
    fmt.Println(aout)
    fmt.Println("the len of pb msg:", len(out))
    jout, _ := json.Marshal(p)
    fmt.Println("the len of json msg:", len(jout))
}
