package route

import (
	"errors"
	"log/slog"
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
func (ih *IssueHandler) PatchIssueByID(ctx echo.Context) error {
	issueID, err := uuid.Parse(ctx.Param("issueID"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseMessage{Status: ERROR, Message: "invalid issueID"})
	}

	req := new(RequestPatchIssue)
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	slog.Info("req", "reqjson", req)

	err = ih.iu.UpdateIssue(issueID, req.Issue.Address, req.Issue.Name, req.Issue.Note, req.Equipments)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "edit issue patch succeeded"})
}

// PutConfirmIssueByID TODO:
// PUT /issue/:issueID
// フォームを確定する
func (ih *IssueHandler) PutConfirmIssueByID(ctx echo.Context) error {
	issueID, err := uuid.Parse(ctx.Param("issueID"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseMessage{Status: ERROR, Message: "invalid issueID"})
	}

	if err := ih.cu.ConfirmIssue(issueID); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success confirm issue"})
}
