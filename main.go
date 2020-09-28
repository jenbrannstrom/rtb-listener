package main

import (
	"google-rtb/config"
	"google-rtb/pkg/logger"
	r "google-rtb/router"
	"log"
	"net/http"
	"time"
)

func main() {
	config.LoadConfig()
	logger.Init()

	router := r.GetRouter()
	s := &http.Server{
		Addr:           r.GetPort(),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.SetKeepAlivesEnabled(false)
	log.Printf("Listening on port %s", r.GetPort())
	log.Fatal(s.ListenAndServe())
}
