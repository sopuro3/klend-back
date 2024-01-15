package repository

import (
	"errors"
)

var (
	// ErrIDIsEmpty  idが空
	ErrIDIsEmpty = errors.New("id is empty")
	ErrEmptyData = errors.New("there is empty data")
)
