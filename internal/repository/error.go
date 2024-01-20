package repository

import (
	"errors"

	"gorm.io/gorm"
)

var (
	// ErrIDIsEmpty  idが空
	ErrIDIsEmpty      = errors.New("id is empty")
	ErrEmptyData      = errors.New("these are empty data")
	ErrRecodeNotFound = gorm.ErrRecordNotFound
	ErrConflict       = errors.New("entity conflicted")
)
