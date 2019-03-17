package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    gorm.Model
    Name string
}

type Profile struct {
    gorm.Model
    UserID int
    User User
    Name string
}

func main() {
    dburl := "root:root@/test_for_gorm?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open("mysql", dburl)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()
    db.AutoMigrate(&User{})
    db.AutoMigrate(&Profile{})
    var user = User{Name:"WinChua"}
    var profile = Profile{Name: "WinChua"}
    db.Create(&user)
    db.Create(&profile)
    db.Model(&user).Related(&profile)
}
