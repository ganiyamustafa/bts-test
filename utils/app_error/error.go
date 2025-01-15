package apperror

import (
	"errors"
	"fmt"

	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AppError struct {
	Err                  error
	CustomHttpStatusCode int
}

func New(message string) *AppError {
	return &AppError{
		Err: errors.New(message),
	}
}

func Sprintf(message string, args ...any) *AppError {
	return &AppError{
		Err: fmt.Errorf(message, args...),
	}
}

func FromError(err error) *AppError {
	if err == nil {
		return nil
	}

	return &AppError{
		Err: err,
	}
}

// get http status code by given error message
func (c *AppError) SetHttpCustomStatusCode(statusCode int) *AppError {
	c.CustomHttpStatusCode = statusCode
	return c
}

// get http status code by given error message
func (c AppError) HttpStatusCode() int {
	// force return status code from app error data
	if c.CustomHttpStatusCode > 0 {
		return c.CustomHttpStatusCode
	}

	// return status code based on error message
	switch {
	case isBadRequestError(c):
		return http.StatusBadRequest
	case isNotFoundError(c):
		return http.StatusNotFound
	case isForbiddenError(c):
		return http.StatusForbidden
	case isUnauthorizedError(c):
		return http.StatusUnauthorized
	case isConflictError(c):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// get http status message by given error message
func (c AppError) HttpStatusMessage() string {
	// force return error if custom status code not zero
	if c.CustomHttpStatusCode > 0 {
		return "error"
	}

	switch {
	case isBadRequestError(c):
		return "bad_request"
	case isNotFoundError(c):
		return "not_found"
	case isForbiddenError(c):
		return "forbidden"
	case isUnauthorizedError(c):
		return "unauthorized"
	case isConflictError(c):
		return "conflict"
	default:
		return "internal_server_error"
	}
}

// get error message
func (c AppError) Error() string {
	if validationErr, ok := c.Err.(validator.ValidationErrors); ok {
		return switchErrorValidation(validationErr)
	}

	return switchGormError(c.Err)
}

// switch and return error by given validation error
func switchErrorValidation(validationErr validator.ValidationErrors) string {
	for _, err := range validationErr {
		field := err.Field()

		// Check Error Type
		switch err.Tag() {
		case "required":
			return fmt.Sprintf("%s is mandatory", field)
		case "email":
			return fmt.Sprintf("%s is not valid email", field)
		case "number":
			return fmt.Sprintf("%s must be numbers only", field)
		case "gte":
			return fmt.Sprintf("%s value must be greater than %s", field, err.Param())
		case "lte":
			return fmt.Sprintf("%s value must be lower than %s", field, err.Param())
		case "min":
			return fmt.Sprintf("%s at least %s characters long", field, err.Param())
		case "max":
			return fmt.Sprintf("the length of %s must be %s characters or fewer", field, err.Param())
		case "startswith":
			return fmt.Sprintf("%s must starts with %s", field, err.Param())
		case "len":
			return fmt.Sprintf("%s length must %s characters", field, err.Param())
		case "oneof":
			return fmt.Sprintf("%s must specify one of %s", field, strings.Join(strings.Split(err.Param(), " "), " or "))
		case "phone":
			return "invalid phone number"
		}
	}

	return validationErr.Error()
}

func switchGormError(err error) string {
	// handle duplicate error
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		// get field constraint
		constraints := strings.Split(err.Error(), "_")
		field := strings.Split(constraints[len(constraints)-1], "\"")[0]

		return fmt.Sprintf("%s already exists", field)
	}

	return err.Error()
}

func isForbiddenError(err AppError) bool {
	switch err.Error() {
	case "forbidden_access":
		return true
	default:
		return false
	}
}

func isUnauthorizedError(err AppError) bool {
	switch err.Error() {
	case "unauthorized":
		return true
	default:
		return false
	}
}

func isNotFoundError(err AppError) bool {
	switch err.Error() {
	case "not_found", gorm.ErrRecordNotFound.Error():
		return true
	default:
		return false
	}
}

func isConflictError(err AppError) bool {
	switch {
	case strings.Contains(err.Error(), "already exists"):
		return true
	}

	switch err.Error() {
	case "conflict":
		return true
	default:
		return false
	}
}

func isBadRequestError(err AppError) bool {
	switch err.Error() {
	case "bad_request":
		return true
	default:
		return false
	}
}
