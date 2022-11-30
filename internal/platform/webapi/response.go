package webapi

const (
	Success = "success"
	Fail    = "fail"
	Err     = "error"
)

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
}
