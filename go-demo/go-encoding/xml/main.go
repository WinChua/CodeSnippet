package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/xml"
)

type Recurlyservers struct {
    XMLName xml.Name `xml:"servers"`
    Version string `xml:"version,attr"`
    Svs []server `xml:"server"`
    //Description string `xml:",innerxml"`
}

type server struct {
    XMLName xml.Name `xml:"server"`
    ServerName string `xml:"serverName"`
    ServerIP string `xml:"serverIP"`
}

func main() {
    if len(os.Args) < 2 {
        fmt.Printf("Usage %s filename\n", os.Args[0])
        return
    }

    f, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error open", os.Args[1],err)
        return
    }
    data, err := ioutil.ReadAll(f)
    if err != nil {
        fmt.Println("Error readall", err)
        return
    }
    fmt.Println(string(data))
    var xmlData Recurlyservers
    err = xml.Unmarshal(data, &xmlData)
    if err != nil {
        fmt.Println("Parse xml:", err)
        return
    }
    fmt.Println(xmlData)

    xmlD, err := xml.Marshal(xmlData)
    if err != nil {
        fmt.Println("Error encoding xml data:", err)
        return
    }
    fmt.Println(string(xmlD))
    fmt.Println("\nEncoding with indent\n")
    xmlD, err = xml.MarshalIndent(xmlData, "  ", "  ")
    if err != nil {
        fmt.Println("Error encoding xml with indent:", err)
        return
    }
    fmt.Println(string(xmlD))
}
