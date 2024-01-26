package model

import "github.com/google/uuid"

type DisplayID struct {
	ID      *uint32 `gorm:"primaryKey,autoIncrement"`
	IssueID uuid.UUID
}
