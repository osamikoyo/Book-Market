package main

import "book-storage/src/api"

func main() {
	s := api.New()
	s.Run()
}