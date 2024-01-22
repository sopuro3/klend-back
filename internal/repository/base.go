//nolint:ireturn // domainとinfraにわけたときにはinterfaceを返す必要がある
package repository

import "gorm.io/gorm"

type BaseRepository interface {
	Atomic(fn func(BaseRepository) error) error
	GetUserRepository() UserRepository
	GetDisplayIDRepository() DisplayIDRepository
	GetIssueRepository() IssueRepository
	GetLoanEntryRepository() LoanEntryRepository
	GetEquipmentRepository() EquipmentRepository
}

type baseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseRepository{
		db: db,
	}
}

func (r *baseRepository) Atomic(fn func(BaseRepository) error) error {
	//nolint:wrapcheck
	return r.db.Transaction(func(tx *gorm.DB) error {
		return fn(NewBaseRepository(tx))
	})
}

func (r *baseRepository) GetUserRepository() UserRepository {
	return NewUserRepository(r.db)
}

func (r *baseRepository) GetDisplayIDRepository() DisplayIDRepository {
	return NewDisplayIDRepository(r.db)
}

func (r *baseRepository) GetIssueRepository() IssueRepository {
	return NewIssueRepository(r.db)
}

func (r *baseRepository) GetLoanEntryRepository() LoanEntryRepository {
	return NewLoanEntryRepository(r.db)
}

func (r *baseRepository) GetEquipmentRepository() EquipmentRepository {
	return NewEquipmentRepository(r.db)
}
