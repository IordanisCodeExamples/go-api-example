package persistence

import "errors"

var (
	// ErrNotFound is the error returned when the requested resource is not found in the database
	ErrNotFound = errors.New("not found")
)
