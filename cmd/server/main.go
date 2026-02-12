package main

import (
	"log"
	"net/http"
	"time"

	"github.com/poojanmishra/SmartLoad/internal/api"
	"github.com/poojanmishra/SmartLoad/internal/solver"
)

func main() {
	optimizer := solver.NewOptimizer()
	handler := api.NewOptimizationHandler(optimizer)
	router := api.NewRouter(handler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Starting SmartLoad Optimization Service on :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
