package model

import "github.com/google/uuid"

type IssueStatus string

const (
	StatusSurvey         IssueStatus = "survey"  // 初期調査
	StatusEquipmentCheck IssueStatus = "check"   // 資機材確認
	StatusConfirm        IssueStatus = "confirm" // 資機材確認(確定)
	StatusReturn         IssueStatus = "return"  // 返却(未納アリ)
	StatusFinish         IssueStatus = "finish"  // 返却完了
)

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
