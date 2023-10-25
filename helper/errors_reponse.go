package helper

import (
	"madyasantosa/ruangkegiatan/dto"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func errorResponse(ctx echo.Context, code int, message string, errors string) error {
	return ctx.JSON(code, dto.ErrorResponse{
		Code:    code,
		Message: message,
		Error:   errors,
	})
}

func StatusNotFound(ctx echo.Context, err error) error {
	return errorResponse(ctx, http.StatusNotFound, err.Error(), "Data Not Found!")
}

func StatusInternalServerError(ctx echo.Context, err error) error {
	logrus.Error(err.Error())
	return errorResponse(ctx, http.StatusInternalServerError, err.Error(), "Internal Server Error!")
}

func StatusBadRequest(ctx echo.Context, err error) error {
	return errorResponse(ctx, http.StatusBadRequest, err.Error(), "Bad Request!")
}
func StatusUnauthorized(ctx echo.Context, err error) error {
	return errorResponse(ctx, http.StatusUnauthorized, err.Error(), "Unauthorized!")
}

func StatusAccountAlreadyExists(ctx echo.Context, err error) error {
	return errorResponse(ctx, http.StatusConflict, err.Error(), "Account Already Exists!")
}
