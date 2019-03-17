package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    if Add(1, 2) != 3 {
        t.FailNow()
    }
}

func TestSub(t *testing.T) {
    if Sub(2, 1) != 1 {
        t.FailNow()
    }
}
