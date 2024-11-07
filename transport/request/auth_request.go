package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// SignUpReq request body for sign up
type SignUpReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (request SignUpReq) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Password, validation.Required, validation.Length(8, 20)),
		validation.Field(&request.Name, validation.Required),
	)
}

// SignInReq request body for sign in
type SignInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request SignInReq) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Password, validation.Required),
	)
}
