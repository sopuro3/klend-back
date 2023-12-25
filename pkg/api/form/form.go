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

type issue struct {
	Address   string    `json:"address"` // 128文字
	Name      string    `json:"name"`    // 128文字
	IssueID   IssueID   `json:"issueId"`
	DisplayID DisplayID `json:"displayId"`
	Note      string    `json:"note"` // 256文字
}

type ResponseForm struct {
	Issue      []issue `json:"issue"`
	TotalIssue int     `json:"totalIssue"`
}

type RequestDeleteForm struct {
	IssueID IssueID `json:"issueId"`
}

// TODO
func GetFormList(ctx echo.Context) error {
	total := 2
	response := ResponseForm{
		Issue: []issue{
			{"小森野1-1-1", "久留米太郎", "018c7765-ffd5-724f-aa7f-227175f54d3f", "0001", "テストデータ"},
			{"浄南町15-3", "久留米次郎", "018c7772-2202-7445-aa24-1bb55e300bdb", "0002", "テストテスト"},
		},
		TotalIssue: total,
	}

	return ctx.JSON(http.StatusOK, response)
}

// TODO
// DELETE /form/[:id]
// フォームを削除
func DeleteForm(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success delete"})
}

// TODO
// POST /form/survey
// フォームを作成
func PostCreateNewSurvey(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success create new survey"})
}
