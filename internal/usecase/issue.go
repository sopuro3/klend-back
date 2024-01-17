package usecase

import (
	"github.com/sopuro3/klend-back/internal/repository"
)

type IssueUseCase struct {
	r repository.BaseRepository
}

func NewIssueUseCase(r repository.BaseRepository) *IssueUseCase {
	return &IssueUseCase{r}
}
