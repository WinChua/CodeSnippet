package main

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func WorkWithUnknowColsV2(db *sql.DB) {
    rows, err := db.Query("SHOW PROCESSLIST")
    if err != nil {
        log.Fatal(err)
    }

    columns, err := rows.ColumnTypes()
    if err != nil {
        log.Fatal(err)
    } else {
        for _, col := range columns {
            fmt.Println(col.Name(), col.ScanType())
        }
    }

    cols, err := rows.Columns()
    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println(cols)
        dest := []interface{} {
            new(uint64),
            new(sql.NullString),
            new(sql.NullString),
            new(sql.NullString),
            new(sql.NullString),
            new(uint32),
            new(sql.NullString),
            new(sql.NullString),
        }
        if len(cols) == 11 {
            fmt.Println("the col is 11")
        } else if len(cols) > 8 {
            fmt.Println("the col is larger 8")
        }

        for rows.Next() {
            err = rows.Scan(dest...)
            if err != nil {
                log.Fatal(err)
            }
            for _, v := range dest {
                switch v.(type) {
                case *uint64:
                    fmt.Println("uint64")
                    fmt.Println(*(v.(*uint64)))
                case *uint32:
                    fmt.Println("uint32")
                    fmt.Println(*(v.(*uint32)))
                default:
                    fmt.Println("string")
                    s := *(v.(*sql.NullString))
                    if s.Valid {
                        fmt.Println(s.String)
                    } else {
                        fmt.Println("<    nil     >")
                    }
                }
            }
        }
    }
}
func WorkWithUnknowColsV1(db *sql.DB) {
    rows, err := db.Query("SHOW PROCESSLIST")
    if err != nil {
        log.Fatal(err)
    }

    columns, err := rows.ColumnTypes()
    if err != nil {
        log.Fatal(err)
    } else {
        for _, col := range columns {
            fmt.Println(col.Name(), col.ScanType())
        }
    }

    cols, err := rows.Columns()
    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println(cols)
        dest := []interface{} {
            new(uint64),
            new(sql.RawBytes),
            new(sql.RawBytes),
            new(sql.RawBytes),
            new(sql.RawBytes),
            new(uint32),
            new(sql.RawBytes),
            new(sql.RawBytes),
        }
        if len(cols) == 11 {
            fmt.Println("the col is 11")
        } else if len(cols) > 8 {
            fmt.Println("the col is larger 8")
        }

        for rows.Next() {
            err = rows.Scan(dest...)
            if err != nil {
                log.Fatal(err)
            }
            for _, v := range dest {
                switch v.(type) {
                case *uint64:
                    fmt.Println("uint64")
                    fmt.Println(*(v.(*uint64)))
                case *uint32:
                    fmt.Println("uint32")
                    fmt.Println(*(v.(*uint32)))
                default:
                    fmt.Println("string")
                    fmt.Println(string(*(v.(*sql.RawBytes))))
                }
            }
        }
    }
}
