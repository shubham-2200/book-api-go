package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *sql.DB

func NewDB(params ...string) *sql.DB {
	var err error
	conString := "root:123456@tcp(127.0.0.1:3306)/store?charset=utf8mb4&parseTime=True&loc=Local"

	log.Print(conString)
	// sql.Open("mysql")
	db, err := sql.Open("mysql", conString)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("used by database", db.Stats().OpenConnections)
	if err != nil {
		fmt.Println("I am at main ")
		fmt.Println("not open", err)
		log.Panic(err)
	}
	DB = db
	return DB
}

func GetDBInstance() *sql.DB {
	return DB
}
