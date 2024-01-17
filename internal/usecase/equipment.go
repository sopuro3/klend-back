package usecase

import (
	"github.com/google/uuid"

	"github.com/sopuro3/klend-back/internal/repository"
)

type EquipmentUseCase struct {
	r repository.BaseRepository
}

func NewEquipmentUseCase(r repository.BaseRepository) *EquipmentUseCase {
	return &EquipmentUseCase{
		r: r,
	}
}

// check issue.isConfirmed TODO
func (eu EquipmentUseCase) CurrentQuantity(equipmentID uuid.UUID) (int32, error) {
	loanEntries, err := eu.r.GetLoanEntryRepository().FindByEquipmentID(equipmentID)
	if err != nil {
		//nolint:wrapcheck
		return 0, err
	}

	var count int32
	for _, loanEntry := range loanEntries {
		count += loanEntry.Quantity
	}

	return count, nil
}
