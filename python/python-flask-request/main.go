package main

import (
    "fmt"
    "net/http"
    "log"
)

func Handle(resp http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(resp, "Hello,World")
}

func main() {
    http.HandleFunc("/", Handle)
    log.Fatal(http.ListenAndServe(":8980", nil))
}

