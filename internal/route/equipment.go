package route

import (
	"errors"
	"log/slog"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/sopuro3/klend-back/internal/model"
	"github.com/sopuro3/klend-back/internal/usecase"
)

type ResponseEquipmentList struct {
	Equipments      []usecase.Equipment `json:"equipments"`
	TotalEquipments int                 `json:"totalEquipments"`
}

type ResponseNewEquipment struct {
	EquipmentID uuid.UUID `json:"id"`
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

		return ctx.JSON(http.StatusInternalServerError, ResponseMessage{ERROR, "Internal Server Error"})
	}

	response := ResponseEquipmentList{
		Equipments:      equipments,
		TotalEquipments: len(equipments),
	}

	return ctx.JSON(http.StatusOK, response)
}

// PostNewEquipment
// POST /equipment
func (eh *EquipmentHandler) PostNewEquipment(c echo.Context) error {
	var equipment usecase.RequestNewEquipment
	if err := c.Bind(&equipment); err != nil {
		slog.Info("bind error", "error", err)

		return c.JSON(http.StatusBadRequest, ResponseMessage{ERROR, "bad request"})
	}

	if err := c.Validate(equipment); err != nil {
		var errs validation.Errors

		errors.As(err, &errs)

		for k, err := range errs {
			slog.Info("validation error", k, err)
		}

		return c.JSON(http.StatusBadRequest, ResponseMessage{ERROR, "bad request"})
	}

	equipmentID, err := eh.eu.CreateNewEquipment(equipment.Name, equipment.Note, equipment.MaxQuantity)
	if err != nil {
		slog.Error("db error", "error", err)

		return c.JSON(http.StatusInternalServerError, ResponseMessage{ERROR, "internal server error"})
	}

	return c.JSON(http.StatusOK, ResponseNewEquipment{equipmentID})
}

// GetEquipmentByID
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

// PutEquipmentByID
// PUT /equipment/[:equipmentId]
func (eh *EquipmentHandler) PutEquipmentByID(c echo.Context) error {
	equipmentID, err := uuid.Parse(c.Param("equipmentID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{ERROR, "invalid equipmentID"})
	}

	var equipment usecase.RequestUpdateEquipment
	if err := c.Bind(&equipment); err != nil {
		slog.Info("bind error", "error", err)
	}

	if err := c.Validate(&equipment); err != nil {
		var errs validation.Errors

		errors.As(err, &errs)

		for k, err := range errs {
			slog.Info("validation error", k, err)
		}

		return c.JSON(http.StatusBadRequest, ResponseMessage{ERROR, "bad request"})
	}

	if err := eh.eu.UpdateEquipment(equipmentID, equipment.Name, equipment.Note, equipment.MaxQuantity); err != nil {
		if errors.Is(err, usecase.ErrRecodeNotFound) {
			return c.JSON(http.StatusBadRequest, ResponseMessage{ERROR, "invalid equipmentID"})
		}

		return err
	}

	return c.JSON(http.StatusOK, ResponseMessage{SUCCESS, "update success"})
}

// DeleteEquipmentByID TODO
// DELETE /equipment/[:equipmentID]
func (eh *EquipmentHandler) DeleteEquipmentByID(c echo.Context) error {
	equipmentID, err := uuid.Parse(c.Param("equipmentID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{ERROR, "invalid equipmentID"})
	}

	if err := eh.eu.DeleteEquipmentByID(equipmentID); err != nil {
		slog.Error("db error", "error", err)

		return c.JSON(http.StatusInternalServerError, ResponseMessage{ERROR, "internal server error"})
	}

	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success delete equipment"})
}
