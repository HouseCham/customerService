package model

type HttpResponse struct {
	Data         interface{} `json:"data"`
	HasError     bool        `json:"hasError"`
	ErrorMessage string      `json:"errorMessage"`
	StatusCode   uint32      `json:"statusCode"`
}
