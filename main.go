package main

import (
	"rep/controller"
	"rep/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initializeRoutes(e *echo.Echo) {
	e.GET("/books", controller.GetBooks)
	e.GET("/", controller.CheckServer)
	e.GET("/bookCount", controller.GetBookCount)
	e.GET("/authorbooks", controller.GetBookByAuthor)
	e.GET("/authors", controller.GetAuthors)
	// e.GET("/books/{bookId}", controller.GetBookByID)

	e.POST("/addTitle", controller.AddTitle)
	e.POST("/addAuthor", controller.StoreAuthor)
	e.POST("/addBook", controller.AddBook)
}
func main() {
	e := echo.New()
	storage.NewDB()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	initializeRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
