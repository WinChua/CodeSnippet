package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Printf("Usage %s ip\n", os.Args[0])
        return
    }
    ip := net.ParseIP(os.Args[1])
    if ip == nil {
        fmt.Println(ip, "is invalid")
    } else {
        fmt.Println("The ip addr is ", ip)
    }

    tcpaddr, err := net.ResolveTCPAddr("tcp", "www.asdf.com:22")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(tcpaddr)
}
