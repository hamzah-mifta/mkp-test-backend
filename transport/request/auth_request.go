package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

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
