package test

import (
	"fmt"
	"madyasantosa/ruangkegiatan/features/users/service"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceImpl_DeleteUser(t *testing.T) {
	t.Run("UserDeletedSuccessfully", func(t *testing.T) {
		userService := &service.UserServiceImpl{
			UserRepository: &MockUsersRepository{
				DeleteUserFunc: func(username string) error {
					return nil
				},
			},
		}

		e := echo.New()

		req := httptest.NewRequest(echo.DELETE, "/your-api-endpoint", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("username")
		c.SetParamValues("testuser")

		err := userService.DeleteUser(c)

		assert.NoError(t, err)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		userService := &service.UserServiceImpl{
			UserRepository: &MockUsersRepository{
				DeleteUserFunc: func(username string) error {
					return fmt.Errorf("User not found")
				},
			},
		}

		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/your-api-endpoint", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("username")
		c.SetParamValues("nonexistentuser")

		err := userService.DeleteUser(c)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "User not found")
	})
}
