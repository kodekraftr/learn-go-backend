package main

import (
	"log"
	"net/http"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) run() error {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}

	log.Printf("Server started at https://127.0.0.1:%s", app.config.addr)

	return server.ListenAndServe()
}
