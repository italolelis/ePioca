package domain

import "errors"

var (
	// ErrNotFound is used when a record is not found
	ErrNotFound = errors.New("record not found")
)
