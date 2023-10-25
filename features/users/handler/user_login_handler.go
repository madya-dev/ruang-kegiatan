package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (uc *UserHandlerImpl) UserLogin(ctx echo.Context) error {
	userLoginRequest := dto.UserLoginRequest{}
	err := ctx.Bind(&userLoginRequest)
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}
	res, err := uc.UserService.UserLogin(ctx, userLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "Invalid username or password") {
			return helper.StatusNotFound(ctx, err)
		}

		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to login"))
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

	return helper.StatusOK(ctx, "Success to login", loginResponse)

}
