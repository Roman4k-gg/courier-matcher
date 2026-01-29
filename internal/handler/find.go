package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"courier-matcher/internal/models"
	"courier-matcher/internal/repository"
)

func FindCouriers(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		latStr := r.URL.Query().Get("lat")
		lonStr := r.URL.Query().Get("lon")
		radiusStr := r.URL.Query().Get("radius")
		vehicle := r.URL.Query().Get("vehicle")
		
		lat, err := strconv.ParseFloat(latStr, 64)
		if err != nil || lat < -90 || lat > 90 {
			http.Error(w, `{"error": "invalid lat"}`, http.StatusBadRequest)
			return
		}
		
		lon, err := strconv.ParseFloat(lonStr, 64)
		if err != nil || lon < -180 || lon > 180 {
			http.Error(w, `{"error": "invalid lon"}`, http.StatusBadRequest)
			return
		}
		
		radius, err := strconv.ParseFloat(radiusStr, 64)
		if err != nil || radius <= 0 {
			http.Error(w, `{"error": "invalid radius"}`, http.StatusBadRequest)
			return
		}
		
		couriers, err := repo.FindNearby(r.Context(), lat, lon, radius, vehicle)
		if err != nil {
			log.Printf("repo error: %v", err)
			http.Error(w, `{"error": "internal error"}`, http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.FindResponse{Couriers: couriers})
	}
}
