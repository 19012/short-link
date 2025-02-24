package storage

import "errors"

// common inforation for all storage implementations

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)
