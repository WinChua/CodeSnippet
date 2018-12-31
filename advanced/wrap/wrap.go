package wrap

import (
    "./fetch"
)

type Subscription interface {
    Updates() <-chan string
    Close() error
}

func Subscribe(fetcher fetch.Fetcher) Subscription {
    s := &sub {
        fetcher: fetcher,
        updates: make(chan string),
        closing: make(chan chan error),
    }
    go s.loop()
    return s
}

