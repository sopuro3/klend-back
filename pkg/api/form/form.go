package form

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sopuro3/klend-back/pkg/api"
)

type (
	IssueID   string // uuidv7
	DisplayID string // 数字4桁
)

type issueStatus string

// TODO 要作りこみ
const (
	IssueStart  issueStatus = "start"
	IssueFinish issueStatus = "finish"
)

type issue struct {
	Address   string      `json:"address"` // 128文字
	Name      string      `json:"name"`    // 128文字
	IssueID   IssueID     `json:"issueId"`
	DisplayID DisplayID   `json:"displayId"`
	Status    issueStatus `json:"status"`
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
	Equipment []struct {
		EquipmentID string `json:"equipmentId"`
		Quantity    int    `json:"quantity"`
	} `json:"equipment"`
}
type ResponseCreateNewIssue struct {
	IssueID string `json:"issueId"`
}

type RequestReturnItem struct {
	Equipment []struct {
		EquipmentID    string `json:"equipmentId"`
		ReturnQuantity int    `json:"returnQuantity"`
	} `json:"equipment"`
}

// GetFormList TODO
func GetFormList(ctx echo.Context) error {
	total := 2
	response := ResponseFormList{
		Issue: []issue{
			{"小森野1-1-1", "久留米太郎", "018c7765-ffd5-724f-aa7f-227175f54d3f", "0001", IssueStart, "テストデータ"},
			{"浄南町15-3", "久留米次郎", "018c7772-2202-7445-aa24-1bb55e300bdb", "0002", IssueFinish, "テストテスト"},
		},
		TotalIssue: total,
	}

	return ctx.JSON(http.StatusOK, response)
}

// DeleteForm TODO
// DELETE /issue/:issueId
// フォームを削除
func DeleteForm(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success delete"})
}

// PostCreateNewSurvey TODO
// POST /issue/survey
// フォームを作成
func PostCreateNewSurvey(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseCreateNewIssue{"018ce372-f35a-705a-9695-ce8ac9c6eff3"})
}

// PostReturnItem TODO
// POST /issue/:issueID/return
// 資機材の返却
func PostReturnItem(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success return item"})
}
