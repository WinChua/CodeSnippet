package main

import (
    "fmt"
    "database/sql"
    "log"
)

func QueryExample(db *sql.DB) {
    rows, err := db.Query("select id, name from users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, name)
    }
}

func PrepareExample(db *sql.DB) {
    stmt, err := db.Prepare("select id, name from users where id = ?")
    if err != nil {
        log.Fatal(err)
    }

    defer stmt.Close()
    rows, err := stmt.Query(13)
    if err != nil {
        log.Fatal(err)
    }

    for rows.Next() {
        err := rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, name)
    }
}

func QueryRowExample(db *sql.DB) {
    err := db.QueryRow("select name from users where id = 1").Scan(&name)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name)
}
