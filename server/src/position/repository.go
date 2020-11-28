package position

import (
	"fmt"
	"src/database"
)

// PositionRepository ....
type PositionRepository struct {
	database *database.Database
}

// NewPositionRepository ...
func NewPositionRepository(db *database.Database) *PositionRepository {
	return &PositionRepository{
		database: db,
	}
}

// GetPositions ...
func (r *PositionRepository) GetPositions() (positions []Position, err error) {
	query := "SELECT id, device_id, timestamp, ST_X(location) as latitude, ST_Y(location) as longitude FROM positions;"

	if err := r.database.Select(&positions, query); err != nil {
		return nil, err
	}

	return positions, nil
}

// AddPosition ...
func (r *PositionRepository) AddPosition(position Position) error {
	query := "INSERT INTO positions (device_id, location) VALUES (:device_id, ST_POINT(:latitude, :longitude));"
	_, err := r.database.NamedQuery(query, position)
	return err
}

// AddPositionByIMEI ...
func (r *PositionRepository) AddPositionByIMEI(model AddPositionRequestModel) error {
	var position Position

	query := "SELECT id FROM positions WHERE timestamp = $1 LIMIT 1;"
	if err := r.database.Get(&position, query, model.Timestamp); err == nil {
		return fmt.Errorf("position with timestamp %s already exists", model.Timestamp)
	}

	query = "INSERT INTO positions (device_id, timestamp, location) VALUES (" +
		"(SELECT id FROM devices WHERE imei = :imei), :timestamp, ST_POINT(:latitude, :longitude))"

	_, err := r.database.NamedQuery(query, model)
	return err
}
