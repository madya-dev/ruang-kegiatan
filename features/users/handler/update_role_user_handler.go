package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (uc *UserHandlerImpl) UpdateRoleUser(ctx echo.Context) error {
	userRoleUpdateRequest := dto.UserRoleUpdateRequest{}
	err := ctx.Bind(&userRoleUpdateRequest)
	fmt.Println(userRoleUpdateRequest)
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}
	res, err := uc.UserService.UpdateRoleUser(ctx, userRoleUpdateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "User not found") {
			return helper.StatusNotFound(ctx, fmt.Errorf("User %s not found!", ctx.Param("username")))
		}

		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed update role user %s", ctx.Param("username")))
	}
	return helper.StatusOK(ctx, "Success to update role user", res)

}
