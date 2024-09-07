package main

import (
	"auth/src/data"
	"auth/src/server"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go data.WorkWithMessage(&wg)

	s := server.New()
	s.Run()
	wg.Wait()
}