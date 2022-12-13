package constants

import "errors"

// Store error
var (
	// ErrKeyNotExists no row
	ErrKeyNotExists = errors.New("key not exists")
	// ErrInvalidDataTypeForValue invalid type for value
	ErrInvalidDataTypeForValue = errors.New("invalid data type for value")
	// ErrInvalidDataTypeForKey invalid type for key
	ErrInvalidDataTypeForKey = errors.New("invalid data type for key")
)
