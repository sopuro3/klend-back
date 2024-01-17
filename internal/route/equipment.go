package route

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/sopuro3/klend-back/internal/model"
	"github.com/sopuro3/klend-back/internal/usecase"
)

type ResponseEquipmentList struct {
	Equipments      []usecase.Equipment `json:"equipments"`
	TotalEquipments int                 `json:"totalEquipments"`
}

type RequestNewEquipment struct {
	Name            string `json:"name"`
	MaxQuantity     int32  `json:"maxQuantity"`
	CurrentQuantity int32  `json:"currentQuantity"`
	Note            string `json:"note"`
}

type ResponseNewEquipment struct {
	EquipmentID string `json:"id"`
}

type EquipmentHandler struct {
	eu *usecase.EquipmentUseCase
}

func NewEquipmentHandler(eu *usecase.EquipmentUseCase) *EquipmentHandler {
	return &EquipmentHandler{
		eu: eu,
	}
}

//nolint:unused
func (eh *EquipmentHandler) modelToResponse(eqModel model.Equipment) (usecase.Equipment, error) {
	currentQuantity, err := eh.eu.CurrentQuantity(eqModel.ID)
	if err != nil {
		//nolint:wrapcheck
		return usecase.Equipment{}, err
	}

	return usecase.Equipment{
		EquipmentID:     eqModel.ID,
		Name:            eqModel.Name,
		CurrentQuantity: currentQuantity,
		MaxQuantity:     eqModel.MaxQuantity,
		Note:            eqModel.Note,
	}, nil
}

// GetEquipmentsList TODO
// GET /equipment
func (eh *EquipmentHandler) GetEquipmentsList(ctx echo.Context) error {
	// TODO
	panic("impl me")

	total := 2
	response := ResponseEquipmentList{
		//nolint:gomnd,lll
		Equipments: []usecase.Equipment{
			{EquipmentID: uuid.MustParse("018c7b9f8c55708f803527a5528e83ed"), Name: "角スコップ", MaxQuantity: 20, CurrentQuantity: 10, Note: "てすとてすとてすと"},
			{EquipmentID: uuid.MustParse("018c7ba8d2df7adcaf3dbe411ce1cb60"), Name: "バケツ", MaxQuantity: 99, CurrentQuantity: 20, Note: "てすとてすとてすと"},
		},
		TotalEquipments: total,
	}

	return ctx.JSON(http.StatusOK, response)
}

// PostNewEquipment TODO
// POST /equipment
func (eh *EquipmentHandler) PostNewEquipment(c echo.Context) error {
	// TODO
	panic("impl me")
	res := ResponseNewEquipment{"018c7b9f8c55708f803527a5528e83ed"}

	return c.JSON(http.StatusOK, res)
}

// GetEquipmentByID TODO
// GET /equipment/[:equipmentId]
func (eh *EquipmentHandler) GetEquipmentByID(ctx echo.Context) error {
	// TODO
	panic("impl me")
	//nolint:gomnd
	res := usecase.Equipment{
		EquipmentID:     uuid.MustParse("018c7b9f8c55708f803527a5528e83ed"),
		Name:            "角スコップ",
		MaxQuantity:     20,
		CurrentQuantity: 10,
		Note:            "てすとてすとてすと",
	}

	return ctx.JSON(http.StatusOK, res)
}

// PutEquipmentByID TODO
// PUT /equipment/[:equipmentId]
func (eh *EquipmentHandler) PutEquipmentByID(c echo.Context) error {
	// TODO
	panic("impl me")
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success update equipment"})
}

// DeleteEquipmentByID TODO
// DELETE /equipment/[:equipmentId]
func (eh *EquipmentHandler) DeleteEquipmentByID(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success delete equipment"})
}
