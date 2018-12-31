package wrap

import (
    "time"
    "./fetch"
)

type sub struct {
    fetcher fetch.Fetcher
    updates chan string
    closed bool
    err error
}

func (s *sub) Updates() <-chan string {
    return s.updates
}

func (s *sub) Close() error {
    s.closed = true
    return s.err
}

func (s *sub) loop() {
    for {
        if s.closed {
            close(s.updates)
            return
        }
        item, next, err := s.fetcher.Fetch()
        if err != nil {
            s.err = err
            time.Sleep(10 * time.Second)
            continue
        }
        s.updates <- item
        time.Sleep(next)
    }
}

