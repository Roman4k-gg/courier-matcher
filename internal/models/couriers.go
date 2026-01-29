package models

type CourierRequest struct {
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Radius  float64 `json:"radius"`
	Vehicle string  `json:"vehicle"`
}

type Courier struct {
	ID      string  `json:"id"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Distance float64 `json:"distance"`
	Rating  float64 `json:"rating"`
	Vehicle string  `json:"vehicle"`
}

type FindResponse struct {
	Couriers []Courier `json:"couriers"`
}
