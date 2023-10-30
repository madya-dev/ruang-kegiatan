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
		if strings.Contains(err.Error(), "User already exists") {
			return helper.StatusAccountAlreadyExists(ctx, err)
		}
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to create user"))
	}

	token, err := helper.GenerateToken(&dto.Token{Username: res.Username, Role: res.Role})
	if err != nil {
		return helper.StatusInternalServerError(ctx, err)
	}

	loginResponse := &dto.UserLoginResponse{
		Username: res.Username,
		Role:     res.Role,
		Token:    token,
	}

	return helper.StatusOK(ctx, "Success to created user", loginResponse)
}
