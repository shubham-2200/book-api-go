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
	books, _ := GetRepBooks()
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
func GetAuthors(c echo.Context) error {
	authors := make([]*model.Author, 0)
	db := storage.GetDBInstance()
	query := "Select Distinct author from shelf"
	row, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// defer db.Close()
	for row.Next() {
		author := new(model.Author)
		err := row.Scan(&author)
		if err != nil {
			fmt.Println(err)
			return err
		}
		authors = append(authors, author)
	}
	return c.JSON(http.StatusOK, authors)
}
