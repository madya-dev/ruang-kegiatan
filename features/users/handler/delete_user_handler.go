package handler

import (
	"fmt"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func (uc *UserHandlerImpl) DeleteUser(ctx echo.Context) error {
	err := uc.UserService.DeleteUser(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "User not found") {
			return helper.StatusNotFound(ctx, fmt.Errorf("User %s not found!", ctx.Param("username")))
		}
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to delete user %s", ctx.Param("username")))
	}

	return helper.StatusNoContent(ctx, "Success to delete data")
}
