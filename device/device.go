package device

type Device struct {
	Name      string `json:"name"`
	IPAddress string `json:"ip_address"`
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
