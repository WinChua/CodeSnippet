package main

import (
    "fmt"
    "net"
    "io"
    "os"
)

func handler(conn net.Conn) {
    target, err := net.Dial("tcp", os.Args[2])
    if err != nil {
        fmt.Println("Dial Error:", err)
        return
    }
    go func() {
        defer target.Close()
        io.Copy(target, conn)
    }()
    go func() {
        defer conn.Close()
        io.Copy(conn, target)
    }()
}

func main() {
    if len(os.Args) != 3 {
        fmt.Printf("Usage %s <source port> <target addr>\n\t%s :4000 mu01:4000\n", os.Args[0], os.Args[0])
        return
    }

    ln, err := net.Listen("tcp", os.Args[1])
    if err != nil {
        fmt.Println("Listen Error:", err)
        return
    }

    defer ln.Close()

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Accept Error:", err)
            return
        }
        go handler(conn)
    }
}
