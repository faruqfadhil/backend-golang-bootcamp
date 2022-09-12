package error

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrBookAlreadyExist = errors.New("book already exist")
	ErrNoRowsAffected   = errors.New("no rows affected")
	ErrUnauthenticated  = errors.New("unauthenticated")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrDatabaseError    = errors.New("database error")
)
