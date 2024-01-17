package usecase

import (
	"github.com/sopuro3/klend-back/internal/repository"
)

type IssueUseCase struct {
	ir repository.IssueRepository
}

func NewIssueUseCase(ir repository.IssueRepository) *IssueUseCase {
	return &IssueUseCase{ir}
}
