package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (uc *UserHandlerImpl) ChangePassword(ctx echo.Context) error {
	authorization := ctx.Request().Header["Authorization"]
	userToken := strings.Split(authorization[0], " ")[1]
	data, err := helper.ExtractToken(userToken)

	changePasswordRequest := dto.ChangePasswordRequest{}
	err = ctx.Bind(&changePasswordRequest)
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}
	err = uc.UserService.ChangePassword(ctx, changePasswordRequest, data.Username)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "User not found") {
			return helper.StatusNotFound(ctx, fmt.Errorf("User %s not found!", data.Username))
		}

		if strings.Contains(err.Error(), "Invalid old password") {
			return helper.StatusNotFound(ctx, err)
		}

		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed change password user %s", data.Username))
	}
	return helper.StatusNoContent(ctx, "Success to change password user")

}
