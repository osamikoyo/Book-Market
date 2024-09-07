package server

import "github.com/labstack/echo/v4"

type Server struct{*echo.Echo}


func New() Server{
	return Server{echo.New()}
}
func(s *Server) Run(){
	
}