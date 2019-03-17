package runner

import (
    //"fmt"
    "os"
    "os/signal"
    "sync"
)

func merge(sigs ...chan os.Signal) chan os.Signal {
    var wg sync.WaitGroup
    out := make(chan os.Signal, 1)
    for _, sig := range sigs {
        go func(sig chan os.Signal) {
            defer wg.Done()
            wg.Add(1)
            for s := range sig {
                out <- s
            }
        }(sig)
    }
    return out
}

func NewSignaler(signals ... os.Signal) chan os.Signal {
    sigs := make([]chan os.Signal, len(signals))
    for _, s := range signals {
        c := make(chan os.Signal, 1)
        signal.Notify(c, s)
        sigs = append(sigs, c)
    }
    return merge(sigs...)
}
