package equipment

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sopuro3/klend-back/pkg/api"
)

type Equipment struct {
	Id              string `json:"id"`
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
	Id string `json:"id"`
}

// TODO
// GET /equipment
func GetEquipmentsList(c echo.Context) error {
	res := ResponseEquipmentList{
		Equipments: []Equipment{
			{Id: "018c7b9f8c55708f803527a5528e83ed", Name: "角スコップ", MaxQuantity: 20, CurrentQuantity: 10, Note: "てすとてすとてすと"},
			{Id: "018c7ba8d2df7adcaf3dbe411ce1cb60", Name: "バケツ", MaxQuantity: 99, CurrentQuantity: 20, Note: "てすとてすとてすと"},
		},
		TotalEquipments: 2,
	}

	return c.JSON(http.StatusOK, res)
}

// TODO
// POST /equipment
func PostNewEquipment(c echo.Context) error {
	res := ResponseNewEquipment{"018c7b9f8c55708f803527a5528e83ed"}
	return c.JSON(http.StatusOK, res)
}

// TODO
// GET /equipment/[:equipmentId]
func GetEquipmentById(c echo.Context) error {
	res := Equipment{Id: "018c7b9f8c55708f803527a5528e83ed", Name: "角スコップ", MaxQuantity: 20, CurrentQuantity: 10, Note: "てすとてすとてすと"}
	return c.JSON(http.StatusOK, res)
}

// TODO
// PUT /equipment/[:equipmentId]
func PutEquipmentById(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success update equipment"})
}

// TODO
// DELETE /equipment/[:equipmentId]
func DeleteEquipmentById(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success delete equipment"})
}
