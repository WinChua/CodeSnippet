package main

import (
    "fmt"
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)
var id int
var name string

func main() {
    db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
    if err != nil {
        log.Fatal(err)
        return
    }
    defer db.Close()

    db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    // fmt.Println("QueryExample(db)")
    // QueryExample(db)
    fmt.Println("PrepareExample(db)")
    PrepareExample(db)
    fmt.Println("QueryRowExample(db)")
    QueryRowExample(db)

    fmt.Println("WorkWithNull(db)")
    WorkWithNull(db)
    fmt.Println("WorkWithUnknowColsV2(db)")
    WorkWithUnknowColsV2(db)
    fmt.Println("WorkWithUnknowColsV1(db)")
    WorkWithUnknowColsV1(db)
    fmt.Println("ErrorHandle(db)")
    ErrorHandle(db)
	fmt.Println("TransationExample(db)")
    ModifyDataWithExecAndQuery(db)
    // never use Query to execute stmt which don't return result

}

