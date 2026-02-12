package api

import (
	"net/http"
)

func NewRouter(h *OptimizationHandler) http.Handler {
	mux := http.NewServeMux()

	// API V1 endpoint
	mux.HandleFunc("/api/v1/load-optimizer/optimize", h.Handle)

	// Health check
	mux.HandleFunc("/healthz", h.HealthCheck)

	return mux
}
