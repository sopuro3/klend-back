package model

import "github.com/google/uuid"

type Issue struct {
	Model
	Address     string       `gorm:"not null"`
	Name        string       `gorm:"not null"`
	Status      string       `gorm:"not null"`
	Note        string       `gorm:"not null"` // 備考がないなら空文字で
	DisplayID   DisplayID    `gorm:"not null"`
	LoanEntries []*LoanEntry `gorm:"not null"`
}

func NewIssue(address, name, status, note string, loanEntries []*LoanEntry) *Issue {
	return &Issue{
		Model: Model{
			ID: uuid.Must(uuid.NewV7()),
		},
		Address:     address,
		Name:        name,
		Status:      status,
		Note:        note,
		LoanEntries: loanEntries,
	}
}
