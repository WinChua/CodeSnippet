package main

import (
    "fmt"
    "encoding/json"
)

type Movie struct {
    Title string
    Year int `json:"released"`
    Color bool `json:"color,omitempty"`
    Actors []string
}

var movies = []Movie{
    {Title: "Casablanca", Year: 1942, Color: false,
        Actors:[]string{"Humphrey Bogart", "Ingrid Bergman"}},
    {Title: "Cool Hand Luke", Year: 1967, Color: true,
        Actors:[]string{"Paul Newman"}},
    {Title: "Bullitt", Year: 1968, Color: true,
        Actors:[]string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
    var data []byte
    data, err := json.Marshal(movies)
    if  err != nil{
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(string(data))
    // if data, err := json.MarshalIndent(movies, "", "  "); err == nil {
    //     fmt.Println(string(data))
    // }
    var titles []struct{Title string}
    if err := json.Unmarshal(data, &titles); err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(titles)

}
