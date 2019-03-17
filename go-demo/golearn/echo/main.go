package main


import (
    "fmt"
    "io"
    "log"
    "net"
)

func main() {
    listener, err := net.Listen("tcp", "localhost:8008")
    if err != nil {
        log.Fatal(err)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err)
            continue
        }

        fmt.Println("connection")
        go handleConn(conn)
    }
}

func handleConn(c net.Conn) {
    defer c.Close()
    io.Copy(c, c)
}
