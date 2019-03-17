package main

import (
    "fmt"
    "errors"
    "strings"
    "time"
    "flag"
)

var flag_example1 = flag.Int("example1", 42, "hello, World")
var flag_gopher string
// var flag_string *string
// var flag_bool *bool

type interval []time.Duration
func (i *interval) String() string {
    return fmt.Sprint(*i)
}
func (i *interval) Set(value string) error {
    if len(*i) > 0 {
        return errors.New("interval flag already set")
    }
    for _, dt := range strings.Split(value, ",") {
        duration, err := time.ParseDuration(dt)
        if err != nil {
            return err
        }
        *i = append(*i, duration)
    }
    return nil
}

var flag_interval interval

func init() {
    const (
        gopher_default = "gopher"
        usage = "default gopher"
    )

    flag.StringVar(&flag_gopher, "gopher", gopher_default, usage)
    flag.StringVar(&flag_gopher, "g", gopher_default, usage+"(shorthand)")
    flag.Var(&flag_interval, "deltaT", "comma-seperated value list")
}

func main() {
    flag.Parse()
    fmt.Println("the flag_int is ", *flag_example1)
    fmt.Println("the flag_string is ", flag_gopher)
    fmt.Println("the flag_bool is ", flag_interval)
}
