package storage

import "errors"

// common inforation for all storage implementations

var (
	ErrUrlNotFound = errors.New("url not found")
	ErrUrlExists   = errors.New("url exists")
)
