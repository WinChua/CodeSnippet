package main

import (
    "fmt"
    "./xkcd"
    "strings"
    "os"
    "flag"
    "io/ioutil"
    "sync"
)

var filename string
var wg sync.WaitGroup

func init () {
    flag.StringVar(&filename, "filename", "", "one number a line.")
}

func main() {
    flag.Parse()
    fmt.Println(flag.Args())
    var num = []string{}
    if len(flag.Args()) == 1 {
        num = append(num, os.Args[1])
    } else {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Errorf("Error Open [%s]: {%s}\n", filename, err)
            return
        }
        num = strings.Split(string(data), "\n")
    }
    var result = []xkcd.XKCD{}
    wg.Add(len(num))
    dataCh := make(chan xkcd.XKCD)
    for _, n := range num {
        go func(n string) {
            defer wg.Done()
            res, err := xkcd.GetURL(n)
            if err != nil {
                fmt.Errorf("Get {%s} Error: {%s}\n", n, err)
                return
            }
            dataCh <- *res
        }(n)
    }
    go func() {
        wg.Wait()
        close(dataCh)
    }()
    for value := range dataCh {
        result = append(result, value)
    }
    for _, v := range result {
        fmt.Println(v)
    }
    fmt.Println("Got", len(result), "results")
}

