package http

import (
	"boilerplate/internal/app/uptime"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func NewHandler() {
	r := mux.NewRouter()
	r.Use()

	r.HandleFunc("/health", uptime.HealthHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
