package models

type HTTPResponse struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}
