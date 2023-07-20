package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseFiles("tmpl/home.html"))

type server struct{}

func NewServer(addr string) *http.Server {
	s := &server{}

	r := mux.NewRouter()
	r.HandleFunc("/", s.handleRoot).Methods("GET")
	r.HandleFunc("/workout", s.handleWorkout).Methods("POST")

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
	templates.ExecuteTemplate(w, "home.html", nil)
}

func (s *server) handleWorkout(w http.ResponseWriter, r *http.Request) {
	d := r.FormValue("description")
	templates.ExecuteTemplate(w, "workout", d)
}
