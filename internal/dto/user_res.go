package dto

type UserRes struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	DeviceId   string `json:"device_id"`
	LastActive string `json:"last_active"`
}
