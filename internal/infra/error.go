package infra

import "fmt"

type apiError struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (e apiError) Error() string {
	return fmt.Sprintf("status code: %d\nmessage: %s", e.StatusCode, e.Message)
}

func NewAPIError(statusCode int, message string) apiError {
	return apiError{
		StatusCode: statusCode,
		Message:    message,
	}
}
