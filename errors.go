package dbx

import "errors"

// package errors
var (
	ErrNotFound           = errors.New("dbx: not found")
	ErrNotSupported       = errors.New("dbx: not supported")
	ErrTableNotSpecified  = errors.New("dbx: table not specified")
	ErrColumnNotSpecified = errors.New("dbx: column not specified")
	ErrInvalidPointer     = errors.New("dbx: attempt to load into an invalid pointer")
	ErrPlaceholderCount   = errors.New("dbx: wrong placeholder count")
	ErrInvalidSliceLength = errors.New("dbx: length of slice is 0. length must be >= 1")
	ErrCantConvertToTime  = errors.New("dbx: can't convert to time.Time")
	ErrInvalidTimestring  = errors.New("dbx: invalid time string")
)
