package route

import (
	"errors"
	"github.com/google/uuid"
	"log/slog"
	"net/http"

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

// GetEquipmentsList
// GET /equipment
func (eh *EquipmentHandler) GetEquipmentsList(ctx echo.Context) error {
	equipments, err := eh.eu.LoadEquipmentList()
	if err != nil {
		slog.Error("GET /equipment", "error", err)
		ctx.JSON(http.StatusInternalServerError, ResponseMessage{ERROR, "Internal Server Error"})
	}

	response := ResponseEquipmentList{
		Equipments:      equipments,
		TotalEquipments: len(equipments),
	}

	return ctx.JSON(http.StatusOK, response)
}

// PostNewEquipment TODO
// POST /equipment
func (eh *EquipmentHandler) PostNewEquipment(c echo.Context) error {
	panic("impl me")
	res := ResponseNewEquipment{"018c7b9f8c55708f803527a5528e83ed"}

	return c.JSON(http.StatusOK, res)
}

// GetEquipmentByID TODO
// GET /equipment/[:equipmentId]
func (eh *EquipmentHandler) GetEquipmentByID(ctx echo.Context) error {
	equipmentID, err := uuid.Parse(ctx.Param("equipmentID"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseMessage{Status: ERROR, Message: "invalid equipmentID"})
	}

	equipment, err := eh.eu.LoadEquipmentByID(equipmentID)
	if err != nil {
		if errors.Is(err, usecase.ErrRecodeNotFound) {
			return ctx.JSON(http.StatusNotFound, ResponseMessage{ERROR, "this equipment is not found"})
		}
		return err
	}

	return ctx.JSON(http.StatusOK, equipment)
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
