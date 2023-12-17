package equipment

import "github.com/labstack/echo/v4"

type IssueId string   //uuidv7
type DisplayId string //数字4桁

type issue struct {
	Adress    string    `json:"adress"`
	Name      string    `json:"name"`
	Id        IssueId   `json:"id"`
	DisplayId DisplayId `json:"displayId"`
	Note      string    `json:"note"`
}

type ResponseForm struct {
	Issue []issue `json:"issue"`
}

type RequestDeleteForm struct {
	Id IssueId `json:"id"`
}

// TODO
func GetFormList(c echo.Context) error {
   testIssue1 := issue{
      Adress: "小森野1-1-1",
      Name: "久留米太郎",
      Id: ,
   }

}
