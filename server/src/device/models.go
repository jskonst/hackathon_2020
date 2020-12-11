package device

// Device ...
type Device struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IMEI      string `json:"imei"`
	AvatarUrl string `json:"avatar_url" db:"avatar_url"`
}
