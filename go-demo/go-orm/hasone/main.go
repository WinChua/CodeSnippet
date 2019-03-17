package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    CreditCard CreditCard
    CreditCardID uint
}

type CreditCard struct {
    gorm.Model
    Number string
}

func main() {
    dburl := "root:root@/test_for_gorm?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open("mysql", dburl)
    if err != nil {
        fmt.Println(err)
        return
    }
    db.AutoMigrate(&CreditCard{})
    db.AutoMigrate(&User{})
    // var card = CreditCard{Number: "9315"}
    // db.Create(&card)
    var user = User{CreditCardID: 0}
    var card CreditCard
    fmt.Println(card)
    db.Model(&user).Related(&card)
    fmt.Println(card)
}
