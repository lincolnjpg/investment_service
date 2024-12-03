package error

import (
	"encoding/json"
	"errors"
	"runtime"
)

type APIError struct {
	Err        error  `json:"err,omitempty"`
	File       string `json:"file,omitempty"`
	Line       int    `json:"line,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
}

func (e APIError) Error() string {
	serializedError, _ := json.Marshal(e.ToMap())

	return string(serializedError)
}

func (e APIError) ToMap() map[string]map[string]interface{} {
	return map[string]map[string]interface{}{
		"error": {
			"message":    e.Err.Error(),
			"file":       e.File,
			"line":       e.Line,
			"statusCode": e.StatusCode,
		},
	}
}

func NewAPIError(message string, statusCode int) APIError {
	_, file, line, _ := runtime.Caller(1)

	return APIError{
		Err:        errors.New(message),
		File:       file,
		Line:       line,
		StatusCode: statusCode,
	}
}
