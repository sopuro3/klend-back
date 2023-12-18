package equipment

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sopuro3/klend-back/pkg/api"
)

type PlannedEquipment struct {
	Equipment
	PlannedQuantity int `json:"plandQuantity"`
}

type ResponseFormData struct {
	Equipments      []PlannedEquipment `json:"equipments"`
	TotalEquipments int                `json:"totalEquipments"`
}

// TODO
// GET /form/[:formId]
// フォームのデータを取得
func GetFormById(c echo.Context) error {
	res := ResponseFormData{
		Equipments: []PlannedEquipment{
			{Equipment{Id: "018c7b9f8c55708f803527a5528e83ed", Name: "角スコップ", MaxQuantity: 20, CurrentQuantity: 10, Note: "てすとてすとてすと"}, 5},
			{Equipment{Id: "018c7ba8d2df7adcaf3dbe411ce1cb60", Name: "バケツ", MaxQuantity: 99, CurrentQuantity: 20, Note: "てすとてすとてすと"}, 10},
		},
		TotalEquipments: 2,
	}

	return c.JSON(http.StatusOK, res)
}

// TODO
// PATCH /form/[:formId]
// フォームを修正
func PatchFormById(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success update planned quantity"})
}

// TODO
// PUT /form/[:formId]/print
// フォームを確定する
func PutConfirmFormById(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success confirm form"})
}