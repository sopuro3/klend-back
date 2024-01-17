package route

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sopuro3/klend-back/internal/usecase"
)

type (
	IssueID   string // uuidv7
	DisplayID string // 数字4桁
)

type IssueStatus string

const (
	StatusSurvey         IssueStatus = "survey"  // 初期調査
	StatusEquipmentCheck IssueStatus = "check"   // 資機材確認
	StatusConfirm        IssueStatus = "confirm" // 資機材確認(確定)
	StatusReturn         IssueStatus = "return"  // 返却(未納アリ)
	StatusFinish         IssueStatus = "finish"  // 返却完了
)

type issue struct {
	Address   string      `json:"address"` // 128文字
	Name      string      `json:"name"`    // 128文字
	IssueID   IssueID     `json:"issueId"`
	DisplayID DisplayID   `json:"displayId"`
	Status    IssueStatus `json:"status"`
	Note      string      `json:"note"` // 256文字
}

type ResponseFormList struct {
	Issue      []issue `json:"issue"`
	TotalIssue int     `json:"totalIssue"`
}

type RequestDeleteForm struct {
	IssueID IssueID `json:"issueId"`
}

type RequestCreateNewIssue struct {
	Issue struct {
		Address string `json:"address"` // 128文字
		Name    string `json:"name"`    // 128文字
		Note    string `json:"note"`    // 256文字
	} `json:"issue"`
	Equipments []struct {
		EquipmentID string `json:"equipmentId"`
		Quantity    int    `json:"quantity"`
	} `json:"equipments"`
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
}

func NewIssueHandler(iu *usecase.IssueUseCase) *IssueHandler {
	return &IssueHandler{iu}
}

// GetFormList TODO
func (ih *IssueHandler) GetFormList(ctx echo.Context) error {
	total := 2
	response := ResponseFormList{
		Issue: []issue{
			{"小森野1-1-1", "久留米太郎", "018c7765-ffd5-724f-aa7f-227175f54d3f", "0001", StatusSurvey, "テストデータ"},
			{"浄南町15-3", "久留米次郎", "018c7772-2202-7445-aa24-1bb55e300bdb", "0002", StatusFinish, "テストテスト"},
		},
		TotalIssue: total,
	}

	return ctx.JSON(http.StatusOK, response)
}

// DeleteForm TODO
// DELETE /issue/:issueId
// フォームを削除
func (ih *IssueHandler) DeleteForm(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success delete"})
}

// PostCreateNewSurvey TODO
// POST /issue/survey
// フォームを作成
func (ih *IssueHandler) PostCreateNewSurvey(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseCreateNewIssue{"018ce372-f35a-705a-9695-ce8ac9c6eff3"})
}

// PostReturnItem TODO
// POST /issue/:issueID/return
// 資機材の返却
func (ih *IssueHandler) PostReturnItem(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success return item"})
}
