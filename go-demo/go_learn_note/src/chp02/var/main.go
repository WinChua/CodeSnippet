package main

import (
    "fmt"
)

func main() {
    var data = 42 // 使用了初始值可以不用声明类型, 由go自行推断类型
    data2 := 23 // 使用 := 进行声明变量, 同时赋值, 可以一次性声明多个变量, 必须包含新变量名称, 只能在函数内部使用
    var data3 int // go将运行期分配的内存初始化为二进制0
    fmt.Println(data, data2)
    data, data2 = data2, data
    fmt.Println(data, data2)
    fmt.Println(data3)
}
