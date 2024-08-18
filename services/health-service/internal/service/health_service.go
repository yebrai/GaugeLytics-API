// services/health-service/internal/service/health_service.go
package service

import (
	"encoding/json"
	"net/http"
)

// HealthCheckHandler handles the health check requests
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "healthy"}
	json.NewEncoder(w).Encode(response)
}
