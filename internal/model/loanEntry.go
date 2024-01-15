package model

import "github.com/google/uuid"

type LoanEntry struct {
	Model
	EquipmentID uuid.UUID `gorm:"not null"`
	Equipment   Equipment
	Quantity    int32     `gorm:"not null"`
	IssueID     uuid.UUID `gorm:"not null"`
}

func NewLoanEntry(quantity int32, equipmentID, issueID uuid.UUID) *LoanEntry {
	return &LoanEntry{
		Model: Model{
			ID: uuid.Must(uuid.NewV7()),
		},
		Quantity:    quantity,
		EquipmentID: equipmentID,
		IssueID:     issueID,
	}
}
