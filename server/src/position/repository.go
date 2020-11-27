package position

import (
	"github.com/jskonst/hackathon_2020/server/database"
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
	err := r.database.Select(&positions, query)
	if err != nil {
		return nil, err
	}

	return positions, nil
}
