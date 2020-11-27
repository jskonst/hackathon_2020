package position

import (
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
