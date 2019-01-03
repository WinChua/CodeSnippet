package main


import (
    "fmt"
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

// Nulable columsn are annoying and lead to a log of ugly code. If you can, avoid them,
// If not, then you'll need to use special types from the datqabase/sql package to handle them,
// or defined your own.

// There are types for nulable booleans, strings, integers, and floats. Here's how you use them:

func WorkWithNull(db *sql.DB) {

    rows, err := db.Query("select name from users where id = 2")
    if err != nil {
        log.Fatal(err)
    }

    for rows.Next() {
        var s sql.NullString
        err := rows.Scan(&s)
        if err != nil {
            log.Fatal(err)
        }
        if s.Valid {
            fmt.Println("The nullable string result is ", s.String)
        } else {
            fmt.Println("The result string is <nil>")
        }
    }
}
