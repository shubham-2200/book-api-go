package main

import (
	"fmt"
	"rep/controller"
	"rep/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// storage
// var DB *gorm.DB

// func NewDB(params ...string) *gorm.DB {
// 	var err error
// 	conString := "root:123456@tcp(127.0.0.1:3306)/store?charset=utf8mb4&parseTime=True&loc=Local"

// 	log.Print(conString)

// 	DB, err = gorm.Open("mysql", conString)
// 	if err != nil {
// 		fmt.Println("I am at main ")
// 		fmt.Println("not open", err)
// 		log.Panic(err)
// 	}

// 	return DB
// }

// func GetDBInstance() *gorm.DB {
// 	return DB
// }

func initializeRoutes(e *echo.Echo) {
	e.GET("/books", controller.GetBooks)
	e.GET("/", controller.CheckServer)
	e.GET("/bookCount", controller.GetBookCount)
	// e.GET("/authorbooks/{authname}", controller.GetBookByAuthor)
	e.GET("/authors", controller.GetAuthors)
	// e.GET("/books/{bookId}", controller.GetBookByID)

	e.POST("/addTitle", controller.AddTitle)
	e.POST("/addAuthor", controller.StoreAuthor)
	e.POST("/addBook", controller.AddBook)
}
func main() {
	e := echo.New()
	storage.NewDB()
	DB := storage.GetDBInstance()
	fmt.Println("used by database", DB.Stats().OpenConnections)
	// db := storage.GetDBInstance()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	initializeRoutes(e)
	// fmt.Println("in main", db.DB().Stats().InUse)
	e.Logger.Fatal(e.Start(":8080"))
}
