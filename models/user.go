package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint
	Name      string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SignupInput struct {
	Name            string `json:"name" validate:"required,min=3,max=20"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=6"`
	Age             int    `json:"age"`
}

type SinginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
