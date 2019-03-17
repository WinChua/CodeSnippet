package xkcd

import (
    "fmt"
    "net/http"
    "encoding/json"
)

const (
    host = "https://xkcd.com/"
    suffix = "/info.0.json"
)

func GetURL(num string) (*XKCD, error) {
    url := host + num + suffix
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("metadata json get error: %s", resp.Status)
    }
    var result XKCD
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }
    return &result, nil
}
