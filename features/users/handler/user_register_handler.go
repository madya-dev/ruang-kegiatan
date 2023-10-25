package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (uc *UserHandlerImpl) RegisterUser(ctx echo.Context) error {
	userCreateRequest := dto.UserCreateRequest{}
	err := ctx.Bind(&userCreateRequest)
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}
	res, err := uc.UserService.CreateUser(ctx, userCreateRequest)

	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to create user"))
	}

	return helper.StatusCreated(ctx, "Success to created user", res)
}
