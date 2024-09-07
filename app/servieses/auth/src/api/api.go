package api

import (
	"auth/src/data"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func GetUserToken(c echo.Context) error{
	var User data.User
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	errbind := c.Bind(&User)
	if errbind != nil {
		loger.Error(errbind.Error())
		return errbind
	}

	user, err := data.GetUserFromToken(User.Token)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}