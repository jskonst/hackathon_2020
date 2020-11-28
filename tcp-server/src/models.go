package main

import "time"

// AddPositionRequestModel ...
type AddPositionRequestModel struct {
	IMEI      string    `json:"imei"`
	Timestamp time.Time `json:"timestamp"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}
