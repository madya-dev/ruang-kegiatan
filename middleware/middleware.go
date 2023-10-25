package middleware

import (
	"fmt"
	"madyasantosa/ruangkegiatan/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header["Authorization"]

		if len(authorization) <= 0 {
			return helper.StatusUnauthorized(c, fmt.Errorf("No token"))
		}

		userToken := strings.Split(authorization[0], " ")

		if len(userToken) <= 1 {
			return helper.StatusBadRequest(c, fmt.Errorf("Invalid token"))
		}

		_, err := helper.ExtractToken(userToken[1])

		if err != nil {
			return helper.StatusBadRequest(c, err)
		}

		return next(c)
	}
}
func AuthMiddlewareAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header["Authorization"]

		if len(authorization) <= 0 {
			return helper.StatusUnauthorized(c, fmt.Errorf("No token"))
		}

		userToken := strings.Split(authorization[0], " ")

		if len(userToken) <= 1 {
			return helper.StatusBadRequest(c, fmt.Errorf("Invalid token"))
		}

		data, err := helper.ExtractToken(userToken[1])

		if err != nil {
			return helper.StatusBadRequest(c, err)
		}
		fmt.Println("roleee", data.Role)
		if data.Role != "admin" {
			return helper.StatusForbidden(c, fmt.Errorf("Access Forbidden!"))
		}

		return next(c)
	}
}
