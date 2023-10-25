package dto

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}
type SuccessResponseWithPagination struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}
type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type SuccessNoDataResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
}
