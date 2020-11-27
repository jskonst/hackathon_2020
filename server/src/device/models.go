package device

// Device ...
type Device struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	IMEI string `json:"imei"`
}