package route

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/sopuro3/klend-back/internal/usecase"
)

type PlannedEquipment struct {
	usecase.Equipment
	PlannedQuantity int `json:"plannedQuantity"`
}

// ResponseFormData
// Issueの取得、更新はこの方を利用して行う
type ResponseFormData struct {
	Issue           issue              `json:"issue"`
	Equipments      []PlannedEquipment `json:"equipments"`
	TotalEquipments int                `json:"totalEquipments"`
}

// GetFormByID TODO
// GET /form/[:formId]
// フォームのデータを取得
func (ih *IssueHandler) GetFormByID(ctx echo.Context) error {
	//nolint:gomnd,lll
	res := ResponseFormData{
		Issue: issue{"小森野1-1-1", "久留米太郎", "018c7765-ffd5-724f-aa7f-227175f54d3f", "0001", StatusSurvey, "テストデータ"},
		//nolint
		Equipments: []PlannedEquipment{
			{usecase.Equipment{EquipmentID: uuid.MustParse("018c7b9f8c55708f803527a5528e83ed"), Name: "角スコップ", MaxQuantity: 20, CurrentQuantity: 10, Note: "てすとてすとてすと"}, 5},
			{usecase.Equipment{EquipmentID: uuid.MustParse("018c7ba8d2df7adcaf3dbe411ce1cb60"), Name: "バケツ", MaxQuantity: 99, CurrentQuantity: 20, Note: "てすとてすとてすと"}, 10},
		},
		TotalEquipments: 2,
	}

	return ctx.JSON(http.StatusOK, res)
}

// PatchIssueByID TODO
// PATCH /issue/:issueID
// フォームを修正
func (ih *IssueHandler) PatchIssueByID(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success update planned quantity"})
}

// PutConfirmIssueByID TODO
// PUT /issue/:issueID
// フォームを確定する
func (ih *IssueHandler) PutConfirmIssueByID(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success confirm issue"})
}
