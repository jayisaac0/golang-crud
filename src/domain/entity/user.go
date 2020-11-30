package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// User entity
type User struct {
	Base
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// UserUpdate struct
type UserUpdate struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
}

// Validate struct
func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Username, validation.Required, validation.Length(4, 50)),
		validation.Field(&user.Email, validation.Required, is.Email, validation.Length(4, 50)),
		validation.Field(&user.Password, validation.Required, validation.Length(4, 50)),
	)
}

// Validate struct
func (updateuser UserUpdate) Validate() error {
	return validation.ValidateStruct(&updateuser,
		validation.Field(&updateuser.Username, validation.Required, validation.Length(4, 50)),
		validation.Field(&updateuser.Email, validation.Required, is.Email, validation.Length(4, 50)),
	)
}
