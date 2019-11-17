package models

// Jwt contains a JWT token.
type Jwt struct {
	Token  string        `json:"token"`
	Custom []interface{} `json:"custom"`
}
