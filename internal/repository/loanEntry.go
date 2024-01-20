//nolint:ireturn // domainとinfraにわけたときにはinterfaceを返す必要がある
package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/sopuro3/klend-back/internal/model"
)

type LoanEntryRepository interface {
	Find(id uuid.UUID) (*model.LoanEntry, error)
	FindByIssueID(issueID uuid.UUID) ([]*model.LoanEntry, error)
	FindByEquipmentID(equipmentID uuid.UUID) ([]*model.LoanEntry, error)
	FindAll() ([]*model.LoanEntry, error)
	Create(equipment *model.LoanEntry) error
	Update(equipment *model.LoanEntry) error
}

type loanEntryRepository struct {
	db *gorm.DB
}

func NewLoanEntryRepository(db *gorm.DB) LoanEntryRepository {
	return &loanEntryRepository{
		db: db,
	}
}

func (lr *loanEntryRepository) Find(id uuid.UUID) (*model.LoanEntry, error) {
	loanEntry := model.LoanEntry{Model: model.Model{ID: id}}

	if err := lr.db.Take(&loanEntry).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &loanEntry, nil
}

func (lr *loanEntryRepository) FindByIssueID(issueID uuid.UUID) ([]*model.LoanEntry, error) {
	var loanEntries []*model.LoanEntry

	if err := lr.db.Where(&model.LoanEntry{IssueID: issueID}).Take(&loanEntries).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return loanEntries, nil
}

func (lr *loanEntryRepository) FindByEquipmentID(equipmentID uuid.UUID) ([]*model.LoanEntry, error) {
	var loanEntries []*model.LoanEntry

	if err := lr.db.Where(&model.LoanEntry{EquipmentID: equipmentID}).Take(&loanEntries).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return loanEntries, nil
}

func (lr *loanEntryRepository) FindAll() ([]*model.LoanEntry, error) {
	var loanEntrys []*model.LoanEntry

	result := lr.db.Find(&loanEntrys)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return loanEntrys, nil
}

func (lr *loanEntryRepository) Create(loanEntry *model.LoanEntry) error {
	if err := lr.db.Create(loanEntry).Error; err != nil {
		return err
	}

	return nil
}

func (lr *loanEntryRepository) Update(loanEntry *model.LoanEntry) error {
	if err := lr.db.Save(loanEntry).Error; err != nil {
		return err
	}

	return nil
}
