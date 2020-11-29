package entity

// User entity
type User struct {
	Base
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
