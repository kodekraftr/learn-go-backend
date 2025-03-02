package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) serverStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is up and working properly"))
}

func main() {
	server := &server{addr: ":8191"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    server.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /", server.serverStatus)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
