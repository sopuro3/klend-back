package model

import "github.com/google/uuid"

type Issue struct {
	Model
	Address     string       `gorm:"not null"`
	Name        string       `gorm:"not null"`
	DisplayID   string       `gorm:"type:char(4)"` // 回収して使い回すしnullable?
	Status      string       `gorm:"not null"`
	Note        string       `gorm:"not null"` // 備考がないなら空文字で
	LoanEntries []*LoanEntry `gorm:"not null"`
}

func NewIssue(address, name, displayID, status, note string, loanEntries []*LoanEntry) *Issue {
	return &Issue{
		Model: Model{
			ID: uuid.Must(uuid.NewV7()),
		},
		Address:     address,
		Name:        name,
		DisplayID:   displayID,
		Status:      status,
		Note:        note,
		LoanEntries: loanEntries,
	}
}
