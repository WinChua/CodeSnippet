package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var c chan os.Signal

func main() {
	c = make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	t := time.After(4 * time.Second)
	future := make(chan interface{})

	go func() {
		time.Sleep(3 * time.Second)
        fmt.Println("Hello")
		future <- 42
	}()

	select {
	case <-c:
		fmt.Println("Interrupt")
	case <-t:
		fmt.Println("Timeout")
	case r := <-future:
		fmt.Println("Result:", r)
	}
    fmt.Printf("%s\n", os.Interrupt)
}

func signaleRegister(sig os.Signal, fn func()) {
    c = make(chan os.Signal, 1)

}
