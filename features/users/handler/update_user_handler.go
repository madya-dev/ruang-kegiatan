package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (uc *UserHandlerImpl) UpdateUser(ctx echo.Context) error {
	authorization := ctx.Request().Header["Authorization"]
	userToken := strings.Split(authorization[0], " ")[1]
	data, _ := helper.ExtractToken(userToken)

	if data.Role != "admin" && data.Username != ctx.Param("username") {
		return helper.StatusForbidden(ctx, fmt.Errorf("Access Forbidden!"))
	}

	userUpdateRequest := dto.UserUpdateRequest{}
	err := ctx.Bind(&userUpdateRequest)
	fmt.Println(userUpdateRequest)
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}
	res, err := uc.UserService.UpdateUser(ctx, userUpdateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "User not found") {
			return helper.StatusNotFound(ctx, fmt.Errorf("User %s Not found!", ctx.Param("username")))
		}

		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed update user %s", ctx.Param("username")))
	}
	return helper.StatusOK(ctx, "Success to update user", res)

}
