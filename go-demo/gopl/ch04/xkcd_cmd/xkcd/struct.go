package xkcd

import "encoding/json"

type XKCD struct {
    Month string
    Num int
    Link string
    Year string
    News string
    SafeTitle string `json:"safe_title"`
    //Transcript string
    Alt string
    Img string
    Title string
    Day string
}


func (x XKCD) String() string {
    data, _ := json.MarshalIndent(x, "", "  ")
    return string(data)
}
