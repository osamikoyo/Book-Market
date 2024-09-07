package api

import (
	"log/slog"
	"net/http"
	"os"
	"register/src/database"

	"github.com/labstack/echo/v4"
)

func Api(){
	
}
func AddUser(c echo.Context) error {
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	var User database.User

	err := c.Bind(&User)
	if err != nil {
		loger.Error(err.Error())
		return err
	}
	err = database.AddUserToDB(User)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Register succes!")
}