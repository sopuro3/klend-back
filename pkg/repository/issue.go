//nolint:ireturn
package repository

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/sopuro3/klend-back/pkg/model"
)

type IssueRepository interface {
	Find(id uuid.UUID) (*model.Issue, error)
	FindAll() ([]*model.Issue, error)
	Create(issue *model.Issue) error
	Update(issue *model.Issue) error
	Delete(issue *model.Issue) error
}

type issueRepository struct {
	db *gorm.DB
}

func NewIssueRepository(db *gorm.DB) IssueRepository {
	return &issueRepository{db}
}

func (ir *issueRepository) Find(id uuid.UUID) (*model.Issue, error) {
	issue := model.Issue{Model: model.Model{ID: id}}

	if err := ir.db.First(&issue).Error; err != nil {
		return nil, err
	}

	return &issue, nil
}

// FindAll 
// レコードが存在しない場合は、len([]*model.Issue)==0を返す
func (ir *issueRepository) FindAll() ([]*model.Issue, error) {
	var issues []*model.Issue

	result := ir.db.Find(&issues)
	if result.Error != nil {
		return issues, result.Error
	}

	return issues, nil
}

func (ir *issueRepository) Create(issue *model.Issue) error {
	if err := ir.db.Create(issue).Error; err != nil {
		return err
	}

	return nil
}

func (ir *issueRepository) Update(issue *model.Issue) error {
	if err := ir.db.Save(issue).Error; err != nil {
		return err
	}

	return nil
}

func (ir *issueRepository) Delete(issue *model.Issue) error {
	if issue.ID == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		//nolint:goerr113,wrapcheck
		return fmt.Errorf("failed delete issue. ID is nil")
	}

	if err := ir.db.Delete(issue).Error; err != nil {
		return err
	}

	return nil
}
