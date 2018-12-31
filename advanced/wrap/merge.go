package wrap


type merge struct {
    subs []Subscription
    updates chan string
    closed bool
    err error
}

func (m *merge) Updates() <-chan string{
    return m.updates
}

func (m *merge) Close() error {
    for _, s := range m.subs {
        s.Close()
    }
    m.closed = true
    return m.err
}

func Merge(subs ...Subscription) Subscription {
    mupdate := make(chan string)
    m := &merge {
        subs: subs,
        updates: mupdate,
    }
    go func() {
        for {
            for _, sub := range subs {
                if m.closed {
                    close(mupdate)
                    return
                }
                select {
                case s := <-(sub.Updates()):
                    mupdate <- s
                default:
                }
            }
        }
    }()
    return m
}
