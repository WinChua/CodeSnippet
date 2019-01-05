package rpcargs

import (
    "fmt"
    "net/rpc"
    "net/http"
)

type Arith int

type Args struct {
    A, B int
}

func (t *Arith) Multiply(a *Args, reply *int) error {
    *reply = a.A * a.B
    return nil
}


func Serve() {
    arith := new(Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()
    err := http.ListenAndServe(":9314", nil)
    if err != nil {
        fmt.Println(err)
    }
}

type Client struct {
    client *rpc.Client
}

func NewClient() (*Client, error) {
    c := new(Client)
    var err error
    c.client, err = rpc.DialHTTP("tcp", ":9314")
    return c, err
}

func (c *Client) Multiply(a int, b int) (int, error) {
    args := Args{A:a, B:b}
    var reply int
    err := c.client.Call("Arith.Multiply", args, &reply)
    return reply, err
}
