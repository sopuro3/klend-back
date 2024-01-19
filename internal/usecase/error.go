package usecase

import (
	"errors"

	"github.com/sopuro3/klend-back/internal/repository"
)

var (
	ErrTooLongString   = errors.New("too long string")
	ErrInvalidQuantity = errors.New("invalid quantity")
	ErrRecodeNotFound  = errors.New("recode not found")
	ErrConflict        = repository.ErrConflict
	ErrIDIsEmpty       = repository.ErrIDIsEmpty
)
