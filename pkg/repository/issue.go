package repository

import (
	"github.com/google/uuid"

	"github.com/sopuro3/klend-back/pkg/model"
)

type IssueRepository interface {
	Find(id uuid.UUID) (*model.Issue, error)
	FindAll() ([]*model.Issue, error)
	Create(issue model.Issue) error
	Update(issue *model.Issue) error
	Delete(issue *model.Issue) error
}
