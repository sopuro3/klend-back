package route

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/sopuro3/klend-back/internal/usecase"
)

// ResponseIssueData
// Issueの取得、更新はこの方を利用して行う
type ResponseIssueData struct {
	Issue           usecase.Issue              `json:"issue"`
	Equipments      []usecase.PlannedEquipment `json:"equipments"`
	TotalEquipments int                        `json:"totalEquipments"`
}

// GetIssueByID TODO
// GET /issue/[:issueID]
// フォームのデータを取得
func (ih *IssueHandler) GetIssueByID(ctx echo.Context) error {
	issueID, err := uuid.Parse(ctx.Param("issueID"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseMessage{Status: ERROR, Message: "invalid issueID"})
	}

	issue, err := ih.iu.GetIssue(issueID)
	if err != nil {
		if errors.Is(err, usecase.ErrRecodeNotFound) {
			return ctx.JSON(http.StatusNotFound, ResponseMessage{ERROR, "this issue is not found"})
		}

		return err
	}

	equipments, err := ih.cu.GetPlannedEquipmentList(issueID)
	if err != nil {
		return err
	}

	response := ResponseIssueData{
		Issue:           issue,
		Equipments:      equipments,
		TotalEquipments: len(equipments),
	}

	return ctx.JSON(http.StatusOK, response)
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
