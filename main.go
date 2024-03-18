package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hellow world"))
	})

	fmt.Println("Server running on port:", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed run the server", err)
	}
}
