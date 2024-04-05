package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong v1.0.1")
	}).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
