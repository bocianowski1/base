package models

type HTTPResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
