package main

import "log"

func main() {
	s := NewServer(":8080")
	log.Fatal(s.ListenAndServe())
}
