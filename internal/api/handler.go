package api

import (
	"encoding/json"
	"net/http"

	"github.com/poojanmishra/SmartLoad/internal/domain"
	"github.com/poojanmishra/SmartLoad/internal/solver"
)

type OptimizationHandler struct {
	Optimizer *solver.Optimizer
}

func NewOptimizationHandler(opt *solver.Optimizer) *OptimizationHandler {
	return &OptimizationHandler{Optimizer: opt}
}

func (h *OptimizationHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.OptimizationRequest
	// Limit request size to prevent DoS
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB limit for safety

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		if err.Error() == "http: request body too large" {
			http.Error(w, "Request body too large", http.StatusRequestEntityTooLarge)
			return
		}
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate basic constraints
	if req.Truck.ID == "" {
		http.Error(w, "Truck ID is required", http.StatusBadRequest)
		return
	}
	// Note: We allow empty orders list, just returns 0 result

	resp := h.Optimizer.Optimize(req)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *OptimizationHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"UP"}`))
}
