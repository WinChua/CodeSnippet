package main
import "fmt"

type embbeded struct  {
    Name string
    Age int
}

func (e *embbeded) EFunc() {
    fmt.Println("Method from embbeded")
}

type Embbeding struct {
    embbeded
    Addr string
}

type EF interface {
    EFunc()
}

func CallEF(a EF) {
    fmt.Println(a)
    a.EFunc()
}

//func (e *Embbeding) EFunc() {
//    fmt.Println("Method from Embedding")
//}

func main() {
    data := Embbeding{}
    fmt.Printf("%+v\n", Embbeding{embbeded{"WinChua", 4}, "Add"})
    data.Addr = "MyWinChua"
    CallEF(&data)

    e := NewEE()
    e.Read()
}
