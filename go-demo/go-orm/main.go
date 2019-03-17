package main

import (
    "fmt"
    "flag"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var dburl string

func init() {
    flag.StringVar(&dburl, "dburl", "root:root@/test_for_gorm?charset=utf8&parseTime=True&loc=Local", "dburl of mysql")
}

type Product struct {
    gorm.Model
    Code string
    Price uint
}

func main() {
    flag.Parse()
    fmt.Println("the dburl is ", dburl)
    //db, err := gorm.Open("sqlite3", "test.db")
    //db, err := gorm.Open("mysql", "root:root@localhost")
    db, err := gorm.Open("mysql",dburl)
    if err != nil {
        fmt.Println(err)
        panic("failed to connect database")
    }
    defer db.Close()

    db.AutoMigrate(&Product{})
    db.Create(&Product{Code: "L9315", Price: 1000})
    var product Product
    db.First(&product, 1)
    db.First(&product, "code = ?", "L9315")
    db.Model(&product).Update("Price", 200)
    db.Delete(&product)
}
