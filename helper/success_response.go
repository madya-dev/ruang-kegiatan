package helper

import (
	"madyasantosa/ruangkegiatan/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func successResponseWithPagination(ctx echo.Context, code int, message string, data any, offset int, limit int, total int) error {
	return ctx.JSON(code, dto.SuccessResponseWithPagination{
		Code:       code,
		Message:    message,
		Data:       data,
		Pagination: dto.Pagination{Offset: offset, Limit: limit, Total: total},
	})
}
func successResponse(ctx echo.Context, code int, message string, data any) error {
	return ctx.JSON(code, dto.SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
func successResponseNoData(ctx echo.Context, code int, message string) error {
	return ctx.JSON(code, dto.SuccessNoDataResponse{
		Code:    code,
		Message: message,
	})
}

func StatusCreated(ctx echo.Context, message string, data any) error {
	return successResponse(ctx, http.StatusCreated, message, data)
}
func StatusCreatedNoContent(ctx echo.Context, message string) error {
	return successResponseNoData(ctx, http.StatusCreated, message)
}
func StatusNoContent(ctx echo.Context, message string) error {
	return successResponseNoData(ctx, http.StatusNoContent, message)
}

func StatusOK(ctx echo.Context, message string, data any) error {
	return successResponse(ctx, http.StatusOK, message, data)
}
func StatusOKWithPagination(ctx echo.Context, message string, data any, offset int, limit int, total int) error {
	return successResponseWithPagination(ctx, http.StatusOK, message, data, offset, limit, total)
}
