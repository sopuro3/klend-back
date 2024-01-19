package usecase

import (
	"errors"
	"github.com/google/uuid"
	"math"
	"unicode/utf8"

	"github.com/sopuro3/klend-back/internal/model"
	"github.com/sopuro3/klend-back/internal/repository"
)

var (
	ErrTooLongString   = errors.New("too long string")
	ErrInvalidQuantity = errors.New("invalid quantity")
	ErrRecodeNotFound  = errors.New("recode not found")
)

const (
	MaxNameLength = 128
	MaxNoteLength = 500
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

func (eu EquipmentUseCase) LoadEquipmentByID(equipmentID uuid.UUID) (*Equipment, error) {
	var equipment *model.Equipment
	var count int32

	err := eu.r.Atomic(func(br repository.BaseRepository) error {
		var err error
		equipment, err = br.GetEquipmentRepository().Find(equipmentID)
		if err != nil {
			return err
		}

		count, err = eu.CurrentQuantity(equipmentID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if equipment == nil {
		return nil, ErrRecodeNotFound
	}

	uEquipment := &Equipment{

		EquipmentID:     equipment.ID,
		Name:            equipment.Name,
		MaxQuantity:     equipment.MaxQuantity,
		CurrentQuantity: count,
		Note:            equipment.Note,
	}

	return uEquipment, nil
}

func (eu EquipmentUseCase) LoadEquipmentList() ([]Equipment, error) {
	equipments := make([]*model.Equipment, 0)
	equipmentsCurrentQty := map[uuid.UUID]int32{}
	loanEntry := make([]*model.LoanEntry, 0)
	err := eu.r.Atomic(func(baseRepository repository.BaseRepository) error {
		er := baseRepository.GetEquipmentRepository()
		lr := baseRepository.GetLoanEntryRepository()

		var err error
		equipments, err = er.FindAll()
		if err != nil {
			return err
		}

		loanEntry, err = lr.FindAll()
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	for _, v := range loanEntry {
		equipmentsCurrentQty[v.EquipmentID] += v.Quantity
	}

	equipmentList := make([]Equipment, 0)

	//nolint:varnamelen
	for _, v := range equipments {
		equipment := Equipment{
			EquipmentID:     v.ID,
			Name:            v.Name,
			MaxQuantity:     v.MaxQuantity,
			Note:            v.Note,
			CurrentQuantity: equipmentsCurrentQty[v.ID],
		}

		equipmentList = append(equipmentList, equipment)
	}

	return equipmentList, nil
}

func (eu EquipmentUseCase) CreateNewEquipment(name, note string, maxQuantity int) (uuid.UUID, error) {
	if utf8.RuneCountInString(name) > MaxNameLength {
		return uuid.UUID{}, ErrTooLongString //nolint:wrapcheck
	}

	if utf8.RuneCountInString(note) > MaxNoteLength {
		return uuid.UUID{}, ErrTooLongString //nolint:wrapcheck
	}

	if maxQuantity < 0 || maxQuantity > math.MaxInt32 {
		return uuid.UUID{}, ErrInvalidQuantity //nolint:wrapcheck
	}

	equipment := model.NewEquipment(name, int32(maxQuantity), note)

	er := eu.r.GetEquipmentRepository()
	if err := er.Create(equipment); err != nil {
		return uuid.UUID{}, err
	}

	return equipment.ID, nil
}
