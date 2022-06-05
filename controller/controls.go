package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"rep/model"
	"rep/storage"

	"github.com/labstack/echo/v4"
)

func StoreAuthor(c echo.Context) error {
	db := storage.GetDBInstance()
	author := new(model.Author)
	if err := c.Bind(&author); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	author.Author = c.QueryParam("author")
	query := "INSERT INTO shelf(author) VALUES(?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Exec(author.Author)
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(result.LastInsertId())
	return c.JSON(http.StatusCreated, author)
}
func AddTitle(c echo.Context) error {
	db := storage.GetDBInstance()
	title := new(model.Title)
	if err := c.Bind(&title); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	title.Title = c.QueryParam("title")
	query := "INSERT INTO shelf(title) VALUES(?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Exec(title.Title)
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(result.LastInsertId())
	return c.JSON(http.StatusCreated, title)
}

func AddBook(c echo.Context) error {
	db := storage.GetDBInstance()
	book := new(model.Book)
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	book.Title = c.QueryParam("title")
	book.Author = c.QueryParam("author")
	query := "INSERT INTO shelf(title,author) VALUES(?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Exec(book.Title, book.Author)
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(result.LastInsertId())
	return c.JSON(http.StatusCreated, book)
}
func CheckServer(c echo.Context) error {
	fmt.Println("Check completed Server running successfully")
	return c.JSON(http.StatusOK, "Running Ok")
}

func GetBooks(c echo.Context) error {
	db := storage.GetDBInstance()
	fmt.Println("connection in use", db.Stats().OpenConnections)
	var books []model.Book
	query := "SELECT DISTINCT title, author FROM shelf"
	row, err := db.Query(query)
	if err != nil {
		fmt.Println("first error:", err)
		return err
	}
	for row.Next() {
		var book model.Book
		var tempTitle, tempAuthor sql.NullString
		if err := row.Scan(&tempTitle, &tempAuthor); err != nil {
			fmt.Println("second error:", err)
			return err
		}
		book.Title = tempTitle.String
		book.Author = tempAuthor.String
		books = append(books, book)
	}
	return c.JSON(http.StatusOK, books)
}
func GetRepBooks() ([]model.Book, error) {
	book := []model.Book{}
	return book, nil
}
func GetBookCount(c echo.Context) error {
	var count model.Count
	db := storage.GetDBInstance()
	query := "Select count(Title) from shelf"
	row := db.QueryRow(query)
	// defer row.Close()
	err := row.Scan(&count.Count)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err == sql.ErrNoRows {
		count.Count = 0
	}

	return c.JSON(http.StatusOK, count)
}
func GetBookByAuthor(c echo.Context) error {
	db := storage.GetDBInstance()
	var books []model.Book
	author := c.QueryParam("author")

	fmt.Println("author:", author)
	query := "SELECT * FROM shelf WHERE author=?"
	row, err := db.Query(query, author)
	if err != nil {
		fmt.Println("first error:", err)
		return err
	}
	for row.Next() {
		var book model.Book
		var tempTitle, tempAuthor sql.NullString
		if err := row.Scan(&tempTitle, &tempAuthor); err != nil {
			fmt.Println("second error:", err)
			return err
		}
		book.Title = tempTitle.String
		book.Author = tempAuthor.String
		books = append(books, book)
	}
	return c.JSON(http.StatusOK, books)

}
func GetAuthors(c echo.Context) error {
	db := storage.GetDBInstance()
	fmt.Println("connection in use", db.Stats().OpenConnections)
	var authors []model.Author
	query := "SELECT DISTINCT author FROM shelf"
	row, err := db.Query(query)
	if err != nil {
		fmt.Println("first error:", err)
		return err
	}
	for row.Next() {
		var author model.Author
		var temp sql.NullString
		if err := row.Scan(&temp); err != nil {
			fmt.Println("second error:", err)
			return err
		}
		author.Author = temp.String
		authors = append(authors, author)
	}
	return c.JSON(http.StatusOK, authors)
}
