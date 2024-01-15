package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/sopuro3/klend-back/internal/model"
	"github.com/sopuro3/klend-back/internal/repository"
)

type Equipment struct {
	EquipmentID     string `json:"equipmentId"`
	Name            string `json:"name"`
	MaxQuantity     int    `json:"maxQuantity"`
	CurrentQuantity int    `json:"currentQuantity"`
	Note            string `json:"note"`
}
type ResponseEquipmentList struct {
	Equipments      []Equipment `json:"equipments"`
	TotalEquipments int         `json:"totalEquipments"`
}

type RequestNewEquipment struct {
	Name            string `json:"name"`
	MaxQuantity     int    `json:"maxQuantity"`
	CurrentQuantity int    `json:"currentQuantity"`
	Note            string `json:"note"`
}

type ResponseNewEquipment struct {
	EquipmentID string `json:"id"`
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

//nolint:unused
func (eu EquipmentUseCase) modelToResponse(eqModel model.Equipment) (Equipment, error) {
	currentQuantity, err := eu.CurrentQuantity(eqModel.ID)
	if err != nil {
		return Equipment{}, err
	}

	return Equipment{
		EquipmentID:     eqModel.ID.String(),
		Name:            eqModel.Name,
		CurrentQuantity: int(currentQuantity),
		MaxQuantity:     int(eqModel.MaxQuantity),
		Note:            eqModel.Note,
	}, nil
}

// TODO: check issue.isConfirmed
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

// GetEquipmentsList TODO
// GET /equipment
func (eu EquipmentUseCase) GetEquipmentsList(ctx echo.Context) error {
	equipments, err := eu.er.FindAll()
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	_ = equipments

	total := 2
	response := ResponseEquipmentList{
		//nolint:gomnd,lll
		Equipments: []Equipment{
			{EquipmentID: "018c7b9f8c55708f803527a5528e83ed", Name: "角スコップ", MaxQuantity: 20, CurrentQuantity: 10, Note: "てすとてすとてすと"},
			{EquipmentID: "018c7ba8d2df7adcaf3dbe411ce1cb60", Name: "バケツ", MaxQuantity: 99, CurrentQuantity: 20, Note: "てすとてすとてすと"},
		},
		TotalEquipments: total,
	}

	return ctx.JSON(http.StatusOK, response)
}

// PostNewEquipment TODO
// POST /equipment
func (eu EquipmentUseCase) PostNewEquipment(c echo.Context) error {
	res := ResponseNewEquipment{"018c7b9f8c55708f803527a5528e83ed"}

	return c.JSON(http.StatusOK, res)
}

// GetEquipmentByID TODO
// GET /equipment/[:equipmentId]
func (eu EquipmentUseCase) GetEquipmentByID(ctx echo.Context) error {
	//nolint:gomnd
	res := Equipment{
		EquipmentID:     "018c7b9f8c55708f803527a5528e83ed",
		Name:            "角スコップ",
		MaxQuantity:     20,
		CurrentQuantity: 10,
		Note:            "てすとてすとてすと",
	}

	return ctx.JSON(http.StatusOK, res)
}

// PutEquipmentByID TODO
// PUT /equipment/[:equipmentId]
func (eu EquipmentUseCase) PutEquipmentByID(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success update equipment"})
}

// DeleteEquipmentByID TODO
// DELETE /equipment/[:equipmentId]
func (eu EquipmentUseCase) DeleteEquipmentByID(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success delete equipment"})
}
