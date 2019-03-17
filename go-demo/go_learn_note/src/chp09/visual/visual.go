package visual

type data struct {
    x int
    Y int
}

func NewData(x, y int) *data {
    return &data{x, y}
}
