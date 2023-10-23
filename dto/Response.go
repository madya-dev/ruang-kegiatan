package dto

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}
type SuccessResponse struct {
	Code       uint32      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}
type SuccessNoDataResponse struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}
type ErrorResponse struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
}
