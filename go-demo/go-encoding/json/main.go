package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
)

type JsonData struct {
    First string `json:""`
    Second string `json:"s"`
    Third string `json:"uP"`
    Forth string `json:"xPf"`
    Fifth string `json:"umMF"`
    Sixth string `json:"cBbdm"`
    Seventh string `json:"CWoSVA"`
    Eighth string `json:"BFKUjzh"`
    Ninth string `json:"hMaqryMR"`
    Tenth string `json:"tzhPgmWIg"`
}

func main() {
    if len(os.Args) < 2 {
        fmt.Printf("Usage %s filename\n", os.Args[0])
        return
    }

    f, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Printf("Open %s Error: %s", os.Args[1], err)
        return
    }
    data, err := ioutil.ReadAll(f)
    if err != nil {
        fmt.Println("ReadAll Error:", err)
        return
    }
    fmt.Println(string(data))

    dataJson := JsonData{}
    err = json.Unmarshal(data, &dataJson)
    if err != nil {
        fmt.Println("Error Parse Json:", err)
        return
    }
    fmt.Println(dataJson)
}
