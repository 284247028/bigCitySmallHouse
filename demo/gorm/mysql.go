package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// User 有多张 CreditCard，UserID 是外键
type User struct {
	gorm.Model
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func main() {
	dsn := "root:#Qiu981016@tcp(127.0.0.1:3306)/house?charset=utf8mb4&parseTime=True&loc=Local"
	conf := mysql.Config{DSN: dsn}
	dial := mysql.New(conf)
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalln(err)
	}
}
