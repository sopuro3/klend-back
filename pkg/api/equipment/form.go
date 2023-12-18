package equipment

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sopuro3/klend-back/pkg/api"
)

type IssueId string   //uuidv7
type DisplayId string //数字4桁

type issue struct {
	Adress    string    `json:"adress"` //128文字
	Name      string    `json:"name"`   // 128文字
	Id        IssueId   `json:"id"`
	DisplayId DisplayId `json:"displayId"`
	Note      string    `json:"note"` // 256文字
}

type ResponseForm struct {
	Issue []issue `json:"issue"`
	Total int     `json:"total"`
}

type RequestDeleteForm struct {
	Id IssueId `json:"id"`
}

// TODO
func GetFormList(c echo.Context) error {
	response := ResponseForm{
		Issue: []issue{
			{"小森野1-1-1", "久留米太郎", "018c7765-ffd5-724f-aa7f-227175f54d3f", "0001", "テストデータ"},
			{"浄南町15-3", "久留米次郎", "018c7772-2202-7445-aa24-1bb55e300bdb", "0002", "テストテスト"},
		},
		Total: 2,
	}

	return c.JSON(http.StatusOK, response)
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
