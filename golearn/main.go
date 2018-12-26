package main
import "fmt"
type Hello struct {
    age int
}

func (a Hello) addone() int {
    a.age += 1
    return a.age
}

func (a *Hello) addonek() int {
    a.age += 1
    return a.age
}

func main() {
    var a Hello
    a.age = 12
    fmt.Println(a.age)
    fmt.Println(a.addone())
    fmt.Println(a.age)
    fmt.Println(a.addonek())
    fmt.Println(a.age)
}
