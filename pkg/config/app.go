package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {

	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	dbName:= os.Getenv("DB_NAME")
	
	dbPassword:= os.Getenv("DB_PASSWORD")
	
	dbUser:= os.Getenv("DB_USER")

	d, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@tcp(127.0.0.1:3306)/"+dbName+"?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("connection failed to open")
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
