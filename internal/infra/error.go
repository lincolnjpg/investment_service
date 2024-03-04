package infra

import (
	"github.com/rotisserie/eris"
)

type APIError struct {
	Err        error `json:"err,omitempty"`
	StatusCode int   `json:"status_code,omitempty"`
}

func (e APIError) Error() string {
	return e.Err.Error()
}

func NewAPIError(message string, statusCode int) APIError {
	return APIError{
		Err:        eris.New(message),
		StatusCode: statusCode,
	}
}
