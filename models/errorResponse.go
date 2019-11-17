package models

// Error contains a HTTP status code and the error message.
type Error struct {
	Code    int    `json:"code"`
	Message string `message:"message"`
}
