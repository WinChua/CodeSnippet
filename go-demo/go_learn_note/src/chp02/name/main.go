package main

import (
    "fmt"
    "./uname" // 相对路径import的参数是路径名称, 路径下面的文件中的package 后面的包名无需与路径名称一致
    // 比如, 在这里, ./uname下面的name.go中的package name, 在使用时候需要用package 后面的名称, name
)

func main() {
    fmt.Println("Hello")
    fmt.Println(name.Hello)
}
