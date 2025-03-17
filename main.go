package main

import (
	"net/http"
	"time"
)

func main() {
	timeoutSeconds := 3

	server := new(http.Server)
	server.Addr = "localhost:8080"
	server.ReadHeaderTimeout = time.Duration(timeoutSeconds) * time.Second

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
