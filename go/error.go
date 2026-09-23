package onet

import "fmt"

// APIError is returned when the O*NET API responds with an HTTP error status.
// StatusCode holds the HTTP status code; Message holds the API's error description.
type APIError struct {
	StatusCode int
	Message    string
}

// Error implements the error interface.
func (e *APIError) Error() string {
	return fmt.Sprintf("onet: HTTP %d: %s", e.StatusCode, e.Message)
}
