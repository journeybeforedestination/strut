package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct{}

func NewServer(addr string) *http.Server {
	s := &server{}

	r := mux.NewRouter()
	r.HandleFunc("/", s.handleRoot).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

func (s *server) handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
