package middleware

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type IValidation interface {
	Validate(data interface{}) error
	Error(err error) error
}

type Validation struct {
	validation *validator.Validate
}

func NewValidation() IValidation {
	return &Validation{
		validation: validator.New(),
	}
}

func (v *Validation) Validate(data interface{}) error {
	err := v.validation.Struct(&data)
	if err != nil {
		return err
	}

	return nil
}

func (v *Validation) Error(err error) error {
	var msg string
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag(){
		case "required":
			msg = fmt.Sprintf("%s is required", e.Field())
		case "min":
			msg = fmt.Sprintf("%s is to short", e.Field())
		case "email":
			msg = fmt.Sprintf("%s data must be email format", e.Field())
		}
	}
}
