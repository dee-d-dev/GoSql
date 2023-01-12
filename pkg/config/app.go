package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:XXXX@tcp(127.0.0.1:3306)/go_rest?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("connection failed to open")
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
