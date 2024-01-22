//nolint:ireturn
package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/sopuro3/klend-back/internal/model"
)

type DisplayIDRepository interface {
	Find(id uint32) (*model.DisplayID, error)
	FindByIssueID(id uuid.UUID) (*model.DisplayID, error)
	Create(issueID uuid.UUID) (*model.DisplayID, error)
}

type displayIDRepository struct {
	db *gorm.DB
}

func NewDisplayIDRepository(db *gorm.DB) DisplayIDRepository {
	return &displayIDRepository{db}
}

func (dir *displayIDRepository) Find(id uint32) (*model.DisplayID, error) {
	displayID := model.DisplayID{ID: &id}

	if err := dir.db.Take(&displayID).Error; err != nil {
		return nil, err
	}

	return &displayID, nil
}

func (dir *displayIDRepository) FindByIssueID(id uuid.UUID) (*model.DisplayID, error) {
	displayID := model.DisplayID{IssueID: id}

	if err := dir.db.Take(&displayID).Error; err != nil {
		return nil, err
	}

	return &displayID, nil
}

func (dir *displayIDRepository) Create(issueID uuid.UUID) (*model.DisplayID, error) {
	displayID := model.DisplayID{
		IssueID: issueID,
	}

	if err := dir.db.Create(&displayID).Error; err != nil {
		return nil, err
	}

	return &displayID, nil
}
