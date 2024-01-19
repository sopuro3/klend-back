package usecase

import (
	"errors"
	"math"
	"unicode/utf8"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"

	"github.com/sopuro3/klend-back/internal/model"
	"github.com/sopuro3/klend-back/internal/repository"
)

const (
	MaxNameLength = 128
	MaxNoteLength = 500
)

type Equipment struct {
	EquipmentID     uuid.UUID `json:"equipmentId"`
	Name            string    `json:"name"`        // utf8で128文字
	MaxQuantity     int32     `json:"maxQuantity"` // qty > 0
	CurrentQuantity int32     `json:"currentQuantity"`
	Note            string    `json:"note"` // utf8で500文字
}

type RequestNewEquipment struct {
	Name        string `json:"name"`
	MaxQuantity int    `json:"maxQuantity"`
	Note        string `json:"note"`
}

func (e RequestNewEquipment) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Name,
			validation.Required.Error("name is required"),
			validation.RuneLength(1, MaxNameLength).Error("name length is 1 ~ 128"),
		),
		validation.Field(&e.MaxQuantity,
			is.Int.Error("current quantity must be int"),
			validation.Min(0),
		),
		validation.Field(&e.Note,
			validation.RuneLength(0, MaxNoteLength).Error("notes are 500 characters max"),
			//			validation.Required.Error("note is required"),
		),
	)
}

type RequestUpdateEquipment = RequestNewEquipment

type EquipmentUseCase struct {
	r repository.BaseRepository
}

func NewEquipmentUseCase(r repository.BaseRepository) *EquipmentUseCase {
	return &EquipmentUseCase{
		r: r,
	}
}

// TODO isConfirm
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

	err := eu.r.Atomic(func(br repository.BaseRepository) error { //nolint:varnamelen
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

func (eu EquipmentUseCase) UpdateEquipment(equipmentID uuid.UUID, name, note string, maxQuantity int) error {
	if utf8.RuneCountInString(name) > MaxNameLength {
		return ErrTooLongString //nolint:wrapcheck
	}

	if utf8.RuneCountInString(note) > MaxNoteLength {
		return ErrTooLongString //nolint:wrapcheck
	}

	if maxQuantity < 0 || maxQuantity > math.MaxInt32 {
		return ErrInvalidQuantity //nolint:wrapcheck
	}

	equipment := &model.Equipment{
		Model:       model.Model{ID: equipmentID},
		Name:        name,
		MaxQuantity: int32(maxQuantity),
		Note:        note,
	}

	if err := eu.r.GetEquipmentRepository().Update(equipment); err != nil {
		if errors.Is(err, repository.ErrRecodeNotFound) {
			return ErrRecodeNotFound
		}

		return err
	}

	return nil
}

func (eu EquipmentUseCase) DeleteEquipmentByID(equipmentID uuid.UUID) error {
	err := eu.r.GetEquipmentRepository().Delete(&model.Equipment{Model: model.Model{ID: equipmentID}})
	if err != nil {
		return err
	}

	err = eu.r.Atomic(func(br repository.BaseRepository) error { //nolint:varnamelen
		loanEntry := &model.LoanEntry{EquipmentID: equipmentID}
		if err := br.GetLoanEntryRepository().Delete(loanEntry); err != nil {
			return err
		}

		equipment := &model.Equipment{Model: model.Model{ID: equipmentID}}

		if err := br.GetEquipmentRepository().Delete(equipment); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
