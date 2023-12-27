package form

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sopuro3/klend-back/pkg/api"
	"github.com/sopuro3/klend-back/pkg/api/equipment"
)

type PlannedEquipment struct {
	equipment.Equipment
	PlannedQuantity int `json:"planedQuantity"`
}

type ResponseFormData struct {
	Issue           issue              `json:"issue"`
	Equipments      []PlannedEquipment `json:"equipments"`
	TotalEquipments int                `json:"totalEquipments"`
}

// GetFormByID TODO
// GET /form/[:formId]
// フォームのデータを取得
func GetFormByID(ctx echo.Context) error {
	res := ResponseFormData{
		Issue: issue{"小森野1-1-1", "久留米太郎", "018c7765-ffd5-724f-aa7f-227175f54d3f", "0001", "テストデータ"},
		//nolint
		Equipments: []PlannedEquipment{

			{equipment.Equipment{EquipmentID: "018c7b9f8c55708f803527a5528e83ed", Name: "角スコップ", MaxQuantity: 20, CurrentQuantity: 10, Note: "てすとてすとてすと"}, 5},
			{equipment.Equipment{EquipmentID: "018c7ba8d2df7adcaf3dbe411ce1cb60", Name: "バケツ", MaxQuantity: 99, CurrentQuantity: 20, Note: "てすとてすとてすと"}, 10},
		},
		//nolint
		TotalEquipments: 2,
	}

	return ctx.JSON(http.StatusOK, res)
}

// PatchFormByID TODO
// PATCH /form/[:formId]
// フォームを修正
func PatchFormByID(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success update planned quantity"})
}

// PutConfirmFormByID TODO
// PUT /form/[:formId]/print
// フォームを確定する
func PutConfirmFormByID(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success confirm form"})
}
