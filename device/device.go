package device

import "github.com/google/uuid"

type Device struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	IPAddress string    `json:"ip_address"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
}
