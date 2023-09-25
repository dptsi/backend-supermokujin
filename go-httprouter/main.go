package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	server := http.Server{
		Addr:    "localhost:8088",
		Handler: router,
	}

	server.ListenAndServe()
}
