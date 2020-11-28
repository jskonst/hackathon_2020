package position

import (
	"time"
)

// Position ...
type Position struct {
	Id        int       `json:"id"`
	DeviceId  int       `json:"device_id" db:"device_id"`
	Timestamp time.Time `json:"timestamp"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}

// AddPositionRequestModel ...
type AddPositionRequestModel struct {
	IMEI      string    `json:"imei"`
	Timestamp time.Time `json:"timestamp"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}
