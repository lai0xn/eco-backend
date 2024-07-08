package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/lai0xn/squid-tech/pkg/types"
)

func NewValidationError(errors validator.ValidationErrors) types.ValidationError {
	validation_err := types.ValidationError{}
	for _, err := range errors {
		e := types.Error{
			NameSpace: err.Namespace(),
			Field:     err.Field(),
			Tag:       err.Field(),
			Kind:      err.Kind().String(),
			Type:      err.Type().String(),
			Value:     err.Value(),
		}
		validation_err.Errors = append(validation_err.Errors, e)
	}
	return validation_err
}
