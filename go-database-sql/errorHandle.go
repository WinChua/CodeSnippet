package main

import (
    "fmt"
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

// We should deal with several type of Error:
// 1. Erros From Iterating ResultSets
//    It's best to check rows.Err() after iterating resultsets.
//    for rows.Next() {
//       ...
//    }
//    if errs := rows.Err(); err != nil { ... }
// 2. Errors From Closing Resultsets
//    Error returned by rows.Close() is the exception to the general rule that
//    we should capture and check for every errors in all database operation.
//    If rows.Close() returns an error, it's unclear what you should do. Logging it
//    may be a good choice or panic it. If that's not sensible, maybe all you can
//    do is just to ignore it.
// 3. Errors From QueryRow()
//    Query may return an empty resultsets, on which a Scan call will return a error
//    sql.ErrNoRows
//    if sql.ErrNoRows happen, just log it, and continue
// 4. Identify Specific Database Errors
// 5. Handling Connection Errors

func ErrorHandle(db *sql.DB) {
    var id int
    var name string
    rows, err := db.Query("select * from users")
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        rows.Scan(&id, &name)
        fmt.Println(id, name)
    }
    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }

    if err = rows.Close(); err != nil {
        log.Println(err)
    }

    err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("no record match id = 1")
            return
        } else {
            log.Fatal(err)
        }
    }
    fmt.Printf("The name match id = 1 is %s\n", name)

}
