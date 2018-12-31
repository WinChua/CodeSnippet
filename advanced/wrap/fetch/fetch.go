package fetch

import (
    "fmt"
    "time"
    "net/http"
    "encoding/json"
)

type Fetcher interface {
    Fetch() (string, time.Duration, error)
}

type Rss struct {
    Url string
}

func Fetch(url string) Fetcher {
    return Rss{Url: url}
}

type respjson struct {
    Title string
    Next time.Duration
}

func (r Rss) Fetch() (string, time.Duration, error) {
    resp, err := http.Get(r.Url)
    if err != nil {
        return "", 0, fmt.Errorf("Get %s error: {%s}", r.Url, err)
    }
    defer resp.Body.Close()
    var data = respjson{}
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return "", 0, fmt.Errorf("Parse resp error: {%s}", err)
    }
    return data.Title, data.Next, nil
}




