package utils

import (
	"encoding/json"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) string {
	if err.Error() == "EOF" {
		return "please check your request body"
	}

	var error string
	switch v := err.(type) {
	case validator.ValidationErrors:
		for _, e := range v {
			if error == "" {
				error = e.Field() + " " + e.Tag()
			} else {
				error = error + ", " + e.Field() + " " + e.Tag()
			}
		}
	case *strconv.NumError, *json.UnmarshalTypeError:
		error = err.Error()
	}

	return error
}
