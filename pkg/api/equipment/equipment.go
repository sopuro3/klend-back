package equipment

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sopuro3/klend-back/pkg/api"
)

type Equipment struct {
	ID              string `json:"id"`
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
	ID string `json:"id"`
}

// TODO
// GET /equipment
func GetEquipmentsList(ctx echo.Context) error {
	total := 2
	response := ResponseEquipmentList{
		Equipments: []Equipment{
			{ID: "018c7b9f8c55708f803527a5528e83ed", Name: "角スコップ", MaxQuantity: 20, CurrentQuantity: 10, Note: "てすとてすとてすと"},
			{ID: "018c7ba8d2df7adcaf3dbe411ce1cb60", Name: "バケツ", MaxQuantity: 99, CurrentQuantity: 20, Note: "てすとてすとてすと"},
		},
		TotalEquipments: total,
	}

	return ctx.JSON(http.StatusOK, response)
}

// TODO
// POST /equipment
func PostNewEquipment(c echo.Context) error {
	res := ResponseNewEquipment{"018c7b9f8c55708f803527a5528e83ed"}

	return c.JSON(http.StatusOK, res)
}

// TODO
// GET /equipment/[:equipmentId]
func GetEquipmentByID(ctx echo.Context) error {
	res := Equipment{
		ID:              "018c7b9f8c55708f803527a5528e83ed",
		Name:            "角スコップ",
		MaxQuantity:     20,
		CurrentQuantity: 10,
		Note:            "てすとてすとてすと",
	}

	return ctx.JSON(http.StatusOK, res)
}

// TODO
// PUT /equipment/[:equipmentId]
func PutEquipmentByID(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success update equipment"})
}

// TODO
// DELETE /equipment/[:equipmentId]
func DeleteEquipmentByID(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success delete equipment"})
}
