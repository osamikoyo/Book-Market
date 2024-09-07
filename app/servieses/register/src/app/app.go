package main

import "register/src/api"

func main() {
	api.Api()
	s:= api.New()
	s.Run()
}