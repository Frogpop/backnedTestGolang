package logger

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

//func ParseJSONReq[T any](r *http.Request, req *T) error {
//	decoder := json.NewDecoder(r.Body)
//	if err := decoder.Decode(&req); err != nil {
//		return fmt.Errorf("Error Decode Request: %w", err)
//	}
//
//	if err := validator.New().Struct(req); err != nil {
//		return fmt.Errorf("Error Validate Request: %w", ValidationErrors(err.(validator.ValidationErrors)))
//	}
//
//	return nil
//}

func ValidationErrors(errs error) error {
	var errorMessages []string
	var ve validator.ValidationErrors
	if errors.As(errs, &ve) {
		for _, err := range ve {
			switch err.ActualTag() {
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf("field %s is required", err.Field()))
			case "min":
				errorMessages = append(errorMessages, fmt.Sprintf("field %s must be greater than %s", err.Field(), err.Param()))
			default:
				errorMessages = append(errorMessages, fmt.Sprintf("field %s is invalid", err.Field()))
			}
		}
		return errors.New(strings.Join(errorMessages, ", "))
	}
	return errs
}
