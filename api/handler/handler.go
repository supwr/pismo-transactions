package handler

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type Validation struct {
	Errors []Field
}

type Field struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func validate(body interface{}) *Validation {
	var validation Validation
	var fields []Field

	v := validator.New(validator.WithRequiredStructEnabled())

	err := v.Struct(body)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fields = append(fields, Field{
				Name:    strings.ToLower(e.Field()),
				Message: "invalid or missing field",
			})
		}
	}

	validation.Errors = fields

	return &validation
}
