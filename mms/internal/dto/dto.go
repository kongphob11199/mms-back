package dto

type StatusResp struct {
	Response Response `json:"response"`
}

type Response string

const (
	OK    Response = "OK"
	ERROR Response = "ERROR"
)
