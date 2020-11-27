package position

import (
	"src/database"
)

// PositionRepository ....
type PositionRepository struct {
	database *database.Database
}

// NewRepository ...
func NewRepository(database *database.Database) *PositionRepository {
	return &PositionRepository{database: database}
}

// GetPositions ...
func (r *PositionRepository) GetPositions() ([]Position, error) {
	var positions []Position

	query := "SELECT id, device_id, timestamp, ST_X(location) as latitude, ST_Y(location) as longitude FROM positions;"

	if err := r.database.Select(&positions, query); err != nil {
		return nil, err
	}

	return positions, nil
}

// AddPosition ...
func (r *PositionRepository) AddPosition(position *Position) error {
	query := "INSERT INTO positions (device_id, location) VALUES (:device_id, ST_POINT(:latitude, :longitude));"

	_, err := r.database.NamedQuery(query, position)
	return err
}
