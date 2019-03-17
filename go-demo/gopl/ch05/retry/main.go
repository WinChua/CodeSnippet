package main

import (
    "flag"
    "log"
    "net/http"
    "fmt"
    "time"
)

func WaitTimeRetry(url string) error {
    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)
    for tries := 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)
        if err == nil {
            return nil
        }
        log.Printf("server not avaliable: %s, retry...", err)
        time.Sleep(time.Second << uint(tries))
    }

    return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

var url string

func init() {
    flag.StringVar(&url, "url", "http://www.baidu.com/", "url")
}

func main() {
    flag.Parse()
    err := WaitTimeRetry(url)
    if err != nil {
        log.Fatalf("Site is down: %v\n", err)
    }
}
