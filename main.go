package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "hello, world")
	})

	log.Fatal(s.ListenAndServe())
}
