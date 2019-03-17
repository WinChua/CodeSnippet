# 在Golang中屏蔽嵌入类型的部分方法

Golang中, 如果一个type中隐式嵌入某一个类型, go的编译器将会认为type实现了嵌入类型实现的接口
当面对某些需要屏蔽嵌入类型部分方法的需求时, 可以通过嵌入子接口以及嵌入类型指针作为field实现

如:

```go

type Interface interface {
    Read()
    Write()
}

type SubInterface interface {
    Read()
}

type Embedded struct{}

//Embedded实现了Interface接口
func (e *Embedded) Read() {}
func (e *Embedded) Write() {}

type Embedding struct {
    SubInterface
    e * Embedded
}

func NewEmbedding() *Embedding {
    e := new(Embedding)
    e.e = new(Embedded)
    e.SubInterface = e.e
    return e
}
```

上面的示例中, Embedding嵌入了Embedded的Read方法, 同时屏蔽了Embedded的Write方法


