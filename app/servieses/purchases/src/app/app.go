package main

import (
	"money/src/server"
)

func main() {
	s := server.New()
	s.Run()
}