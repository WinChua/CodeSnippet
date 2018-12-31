package wrap

import (
    "fmt"
    "time"
    "./fetch"
)

type sub struct {
    fetcher fetch.Fetcher
    updates chan string
    err error
    closing chan chan error
}

func (s *sub) Updates() <-chan string {
    return s.updates
}

func (s *sub) Close() error {
    errc := make(chan error)
    //s.closing is chan chan error. such struct is called require responsed.
    s.closing <- errc
    return <-errc
}

func (s *sub) loop() {
    var pending []string
    var next time.Duration
    var err error
    for {
        var first string
        var update chan string
        if len(pending) > 0 {
            first = pending[0]
            update = s.updates
        }
        startFetch := time.After(next)
        select {
        // you should initialization s.closing in Subscribe function
        case errc := <-s.closing:
            errc <- err
            close(s.updates)
            return
        case <-startFetch:
            var item string
            item, next, err = s.fetcher.Fetch()
            if err == nil {
                pending = append(pending, item)
            }
        case update <- first:
            pending = pending[1:]
        }
    }
}

