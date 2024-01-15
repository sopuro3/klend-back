package api

type Status string

const (
	SUCCESS Status = "success"
	ERROR   Status = "error"
)

type ResponseMessage struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
}
