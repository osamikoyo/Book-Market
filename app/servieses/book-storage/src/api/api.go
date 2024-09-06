package api

import (
	"book-storage/src/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBooks(c echo.Context) error {
	books, err := database.GetBooksFromDB()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, books)
}