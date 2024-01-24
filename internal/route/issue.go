package route

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sopuro3/klend-back/internal/usecase"
)

type ResponseFormList struct {
	Issue      []usecase.Issue `json:"issue"`
	TotalIssue int             `json:"totalIssue"`
}

type RequestDeleteForm struct {
	IssueID usecase.IssueID `json:"issueId"`
}

type RequestCreateNewIssue struct {
	Issue struct {
		Address string `json:"address"` // 128文字
		Name    string `json:"name"`    // 128文字
		Note    string `json:"note"`    // 256文字
	} `json:"issue"`
	Equipments []usecase.EquipmentWithQuantity `json:"equipments"`
}
type ResponseCreateNewIssue struct {
	IssueID string `json:"issueId"`
}

type RequestReturnItem struct {
	Equipments []struct {
		EquipmentID    string `json:"equipmentId"`
		ReturnQuantity int    `json:"returnQuantity"`
	} `json:"equipments"`
}

type IssueHandler struct {
	iu *usecase.IssueUseCase
	cu *usecase.ConfirmUsecase
}

func NewIssueHandler(iu *usecase.IssueUseCase, cu *usecase.ConfirmUsecase) *IssueHandler {
	return &IssueHandler{iu, cu}
}

// GetIssueList TODO
func (ih *IssueHandler) GetIssueList(ctx echo.Context) error {
	issues, err := ih.iu.GetIssueAll()
	if err != nil {
		return err
	}

	response := ResponseFormList{
		Issue:      issues,
		TotalIssue: len(issues),
	}

	return ctx.JSON(http.StatusOK, response)
}

// DeleteForm TODO
// DELETE /issue/:issueId
// フォームを削除
func (ih *IssueHandler) DeleteForm(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success delete"})
}

// PostCreateNewIssue TODO currentQuantity check
// POST /issue/survey
// フォームを作成
func (ih *IssueHandler) PostCreateNewIssue(ctx echo.Context) error {
	req := new(RequestCreateNewIssue)
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	id, err := ih.iu.CreateIssue(req.Issue.Address, req.Issue.Name, req.Issue.Note, req.Equipments)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, ResponseCreateNewIssue{id.String()})
}

// PostReturnItem TODO
// POST /issue/:issueID/return
// 資機材の返却
func (ih *IssueHandler) PostReturnItem(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success return item"})
}
