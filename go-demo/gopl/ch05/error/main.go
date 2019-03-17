package main

import "fmt"

func getError(name string) error{
    return fmt.Errorf(name)
}

func main() {
    fmt.Println(getError("Hello,World"))
}
