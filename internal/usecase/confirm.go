package usecase

import (
	"github.com/google/uuid"

	"github.com/sopuro3/klend-back/internal/model"
	"github.com/sopuro3/klend-back/internal/repository"
)

type PlannedEquipment struct {
	Equipment
	PlannedQuantity int32 `json:"plannedQuantity"`
}

type ConfirmUsecase struct {
	r repository.BaseRepository
}

func NewConfirmUseCase(r repository.BaseRepository) *ConfirmUsecase {
	return &ConfirmUsecase{
		r: r,
	}
}

func (cu *ConfirmUsecase) GetPlannedEquipmentList(issueID uuid.UUID) ([]PlannedEquipment, error) {
	equipments := make([]*model.Equipment, 0)
	equipmentsCurrentQty := map[uuid.UUID]int32{}
	equipmentsPlannedQty := map[uuid.UUID]int32{}
	loanEntry := make([]*model.LoanEntry, 0)

	err := cu.r.Atomic(func(br repository.BaseRepository) error { //nolint:varnamelen
		var err error
		equipments, err = br.GetEquipmentRepository().FindAll()
		if err != nil {
			return err
		}

		loanEntry, err = br.GetLoanEntryRepository().FindAll()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// check issue.Status FIXME
	for _, v := range loanEntry {
		if v.IssueID == issueID {
			equipmentsPlannedQty[v.EquipmentID] += v.Quantity
		} else {
			equipmentsCurrentQty[v.EquipmentID] += v.Quantity
		}
	}

	plannedEquipmentList := make([]PlannedEquipment, 0, len(equipments))

	//nolint:varnamelen
	for _, v := range equipments {
		plannedEquipment := PlannedEquipment{
			Equipment: Equipment{
				EquipmentID:     v.ID,
				Name:            v.Name,
				MaxQuantity:     v.MaxQuantity,
				Note:            v.Note,
				CurrentQuantity: equipmentsCurrentQty[v.ID],
			},
			PlannedQuantity: equipmentsPlannedQty[v.ID],
		}

		plannedEquipmentList = append(plannedEquipmentList, plannedEquipment)
	}

	return plannedEquipmentList, nil
}

func (cu *ConfirmUsecase) ConfirmIssue(issueID uuid.UUID) error {
	err := cu.r.Atomic(func(br repository.BaseRepository) error {
		ir := br.GetIssueRepository() //nolint:varnamelen

		var err error

		issue, err := ir.Find(issueID)
		if err != nil {
			return ErrRecodeNotFound
		}

		if issue.Status != string(model.StatusEquipmentCheck) {
			return ErrInvalidStatus
		}

		err = ir.Update(&model.Issue{
			Model:  model.Model{ID: issueID},
			Status: string(model.StatusConfirm),
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
