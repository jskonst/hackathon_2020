package position

import (
	"time"
)

// Position ...
type Position struct {
	Id        int       `json:"id"`
	DeviceId  int       `json:"device_id" db:"device_id"`
	Timestamp time.Time `json:"timestamp"`
	Latitude  float32   `json:"latitude"`
	Longitude float32   `json:"longitude"`
}
