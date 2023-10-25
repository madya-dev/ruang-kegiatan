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
		userToken := strings.Split(authorization[0], " ")[1]
		_, err := helper.ExtractToken(userToken)

		if err != nil {
			return helper.StatusBadRequest(c, err)
		}

		return next(c)
	}
}
func AuthMiddlewareAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authorization := c.Request().Header["Authorization"]
		userToken := strings.Split(authorization[0], " ")[1]
		data, err := helper.ExtractToken(userToken)

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
