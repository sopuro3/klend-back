package usecase

import (
	"github.com/google/uuid"
	"github.com/sopuro3/klend-back/internal/model"

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
	r repository.BaseRepository
}

func NewEquipmentUseCase(r repository.BaseRepository) *EquipmentUseCase {
	return &EquipmentUseCase{
		r: r,
	}
}

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

func (eu EquipmentUseCase) EquipmentList() ([]Equipment, error) {
	equipments := make([]*model.Equipment, 0)
	eCurrentQty := map[uuid.UUID]int32{}
	err := eu.r.Atomic(func(baseRepository repository.BaseRepository) error {
		er := baseRepository.GetEquipmentRepository()
		lr := baseRepository.GetLoanEntryRepository()

		var err error
		equipments, err = er.FindAll()
		if err != nil {
			return err
		}

		loanEntrys, err := lr.FindAll()
		if err != nil {
			return err
		}

		for _, v := range loanEntrys {
			eCurrentQty[v.EquipmentID] += v.Quantity
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	equipmentList := make([]Equipment, 0)

	for _, v := range equipments {
		equipment := Equipment{
			EquipmentID:     v.ID,
			Name:            v.Name,
			MaxQuantity:     v.MaxQuantity,
			Note:            v.Note,
			CurrentQuantity: eCurrentQty[v.ID],
		}

		equipmentList = append(equipmentList, equipment)
	}

	return equipmentList, nil
}
