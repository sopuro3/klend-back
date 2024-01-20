package usecase

import "errors"

var (
	ErrTooLongString   = errors.New("too long string")
	ErrInvalidQuantity = errors.New("invalid quantity")
	ErrRecodeNotFound  = errors.New("recode not found")
)
