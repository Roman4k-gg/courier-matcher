package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math"
	
	"courier-matcher/internal/models"
)

type Repository interface {
	FindNearby(ctx context.Context, lat, lon, radius float64, vehicle string) ([]models.Courier, error)
}

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) FindNearby(ctx context.Context, lat, lon, radius float64, vehicle string) ([]models.Courier, error) {
	query := `
		SELECT id::text, lat, lon, rating, vehicle 
		FROM couriers 
		WHERE status = 'free' 
		AND (vehicle = $1 OR $1 = '')
		AND ABS(lat - $2) < 0.05  -- ±3км широты
		AND ABS(lon - $3) < 0.05  -- ±3км долготы
		ORDER BY ABS(lat - $2) + ABS(lon - $3)
		LIMIT 10
	`
	rows, err := r.db.QueryContext(ctx, query, vehicle, lat, lon)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()
	
	var couriers []models.Courier
	for rows.Next() {
		var c models.Courier
		if err := rows.Scan(&c.ID, &c.Lat, &c.Lon, &c.Rating, &c.Vehicle); err != nil {
			log.Printf("scan error: %v", err)
			continue
		}
		c.Distance = haversine(lat, lon, c.Lat, c.Lon)
		couriers = append(couriers, c)
	}
	
	return couriers, nil
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180
	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + 
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1Rad)*math.Cos(lat2Rad)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
