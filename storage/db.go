package storage

import (
	"fmt"
	"log"
	"rep/model"

	"gorm.io/driver/mysql"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func NewDB(params ...string) {

	conString := "root:123456@tcp(127.0.0.1:3306)/store?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(conString)
	log.Print(conString)
	// sql.Open("mysql")
	DB, err = gorm.Open(mysql.Open(conString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("used by database", db.)
	db, err := DB.DB()
	if err != nil {
		fmt.Println("I am at main ")
		fmt.Println("not open", err)
		log.Panic(err)
	}
	fmt.Println("Database connected successfully ")
	fmt.Println("connection in use", db.Stats().OpenConnections)
	DB.AutoMigrate(&model.Book{})
	fmt.Println("connection in use", db.Stats().OpenConnections)
	// fmt.Println("connection in use", DB.Statement.Config)
}

func GetDBInstance() *gorm.DB {

	fmt.Println("in getDBInstance")
	db, err := DB.DB()
	fmt.Println("connection in use", db.Stats().OpenConnections)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connection in use", db.Stats().OpenConnections)

	return DB
}
