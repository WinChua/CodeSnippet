package  main

import (
    "os"
    "fmt"
    "bytes"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Errorf("Usage %s num", os.Args[0])
        return 
    }
    var buf bytes.Buffer
    var i = 0
    var j = len(os.Args[1])
    if os.Args[1][i] == '-' || os.Args[1][i] == '+' {
        buf.WriteByte(os.Args[1][i])
        i += 1
    }
    if (j - i) % 3 != 0 {
        buf.WriteString(os.Args[1][i:i+(j-i)%3])
        i += (j-i)%3
    } else {
        buf.WriteString(os.Args[1][i:i+3])
        i += 3
    }
    
    for ; i < j; i += 3 {
        buf.WriteString(","+os.Args[1][i:i+3])
    }

    fmt.Println(buf.String())
}


