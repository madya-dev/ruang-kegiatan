package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (uc *UserHandlerImpl) GetUserByUsername(ctx echo.Context) error {
	res, err := uc.UserService.GetUserByUsername(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return helper.StatusNotFound(ctx, fmt.Errorf("User %s Not Found!", ctx.Param("username")))
		}
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to Get User %s", ctx.Param("username")))
	}

	return helper.StatusOK(ctx, "Success to get data", res)
}
