package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct{
	*echo.Echo
}
func New() Server {
	return Server{echo.New()}
}
func (s Server) Run(){
	s.Use(middleware.Logger())


	s.GET("/books", GetBooks)
	s.Logger.Panic(s.Start(":2021"))
}