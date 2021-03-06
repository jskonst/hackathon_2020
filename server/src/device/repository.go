package device

import (
	"src/database"
)

// DeviceRepository ...
type DeviceRepository struct {
	database *database.Database
}

// NewDeviceRepository ...
func NewDeviceRepository(db *database.Database) *DeviceRepository {
	return &DeviceRepository{
		database: db,
	}
}

// GetDevices ...
func (r *DeviceRepository) GetDevices() (devices []Device, err error) {
	query := "SELECT * FROM devices;"

	if err := r.database.Select(&devices, query); err != nil {
		return nil, err
	}

	return devices, nil
}

// GetDeviceByIMEI ...
func (r *DeviceRepository) GetDeviceByIMEI(imei string) (device Device, err error) {
	query := "SELECT * FROM devices WHERE imei = $1 LIMIT 1;"

	if err := r.database.QueryRowx(query, imei).StructScan(&device); err != nil {
		return device, err
	}

	return device, nil
}

// AddDevice ...
func (r *DeviceRepository) AddDevice(device Device) error {
	query := "INSERT INTO devices (name, imei) VALUES (:name, :imei);"
	_, err := r.database.NamedQuery(query, device)
	return err
}