package entity

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrUnexpected          = errors.New("unexpected error")
	ErrInternalServerError = errors.New("internal server error")
)
