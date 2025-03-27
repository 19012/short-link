package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url not round")
	ErrURLExists   = errors.New("url exists")
)
