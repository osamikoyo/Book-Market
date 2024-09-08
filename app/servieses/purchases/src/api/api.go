package api

import (
	"log/slog"
	"money/src/database"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func GetPurshes(c echo.Context) error {
	var u database.User
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := c.Bind(&u)
	if err != nil {
		loger.Error(err.Error())
		return err
	}

	data, errfromdb := database.GetPursh(u)
	if errfromdb != nil {
		loger.Error(errfromdb.Error())
		return errfromdb
	}
	return c.JSON(http.StatusOK, data)
}