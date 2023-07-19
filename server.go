package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type server struct{}

func NewServer(addr string) *http.Server {
	s := &server{}

	r := mux.NewRouter()
	r.HandleFunc("/", s.handleRoot).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: withLogging(r),
	}
}

func withLogging(h http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		uri := r.RequestURI
		method := r.Method
		h.ServeHTTP(w, r)

		duration := time.Since(start)

		log.Printf("%s\t%s\t%s", duration, method, uri)
	}
	return http.HandlerFunc(logFn)
}

func (s *server) handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
