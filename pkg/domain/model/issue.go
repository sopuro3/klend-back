package model

import "github.com/google/uuid"

type Issue struct {
	Model
	Address     string `gorm:"not null"`
	Name        string `gorm:"not null"`
	DisplayID   string `gorm:"type:char(4)"` // 回収して使い回すしnullable?
	Status      string `gorm:"not null"`
	Note        string `gorm:"not null"` // 備考がないなら空文字で
	IsConfirmed bool   `gorm:"not null"`
	Loan
}

type Loan struct {
	LoanEntries []LoanEntry `gorm:"not null"`
}

type LoanEntry struct {
	Model
	EquipmentID uuid.UUID `gorm:"not null"`
	Quantity    int32     `gorm:"not null"`
	IssueID     uuid.UUID `gorm:"not null"`
}
