package usecase

import (
	"github.com/google/uuid"

	"github.com/sopuro3/klend-back/internal/repository"
)

type Equipment struct {
	EquipmentID     uuid.UUID `json:"equipmentId"`
	Name            string    `json:"name"`
	MaxQuantity     int32     `json:"maxQuantity"`
	CurrentQuantity int32     `json:"currentQuantity"`
	Note            string    `json:"note"`
}

type EquipmentUseCase struct {
	er repository.EquipmentRepository
	lr repository.LoanEntryRepository
}

func NewEquipmentUseCase(er repository.EquipmentRepository, lr repository.LoanEntryRepository) *EquipmentUseCase {
	return &EquipmentUseCase{
		er: er,
		lr: lr,
	}
}

func (eu EquipmentUseCase) CurrentQuantity(equipmentID uuid.UUID) (int32, error) {
	loanEntries, err := eu.lr.FindByEquipmentID(equipmentID)
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

func (eu EquipmentUseCase) EquipmentList() ([]Equipment, error) {
	equipments, err := eu.er.FindAll()
	if err != nil {
		return nil, err
	}

	var equipmentList []Equipment
	for _, v := range equipments {
		equipment := Equipment{
			EquipmentID: v.ID,
			Name:        v.Name,
			MaxQuantity: v.MaxQuantity,
			Note:        v.Note,
		}

		equipmentList = append(equipmentList, equipment)
	}

	return equipmentList, nil

}
