package test

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/users/service"
	"madyasantosa/ruangkegiatan/model"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceImpl_UpdateRoleUser(t *testing.T) {
	userService := &service.UserServiceImpl{
		UserRepository: &MockUsersRepository{
			GetUserByUsernameFunc: func(username string) (*model.User, error) {
				if username == "testuser" {
					return &model.User{
						Username: "testuser",
						Role:     "user",
					}, nil
				}
				return nil, fmt.Errorf("User not found")
			},
			UpdateRoleUserFunc: func(user *model.User, username string) (*model.User, error) {
				return user, nil
			},
		},
		Validate: validator.New(),
	}

	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/your-api-endpoint", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("username")
	c.SetParamValues("testuser")

	t.Run("ValidUser", func(t *testing.T) {
		roleUpdateRequest := dto.UserRoleUpdateRequest{
			Role: "admin",
		}

		userResponse, err := userService.UpdateRoleUser(c, roleUpdateRequest)

		assert.NoError(t, err)
		assert.NotNil(t, userResponse)
		assert.Equal(t, "admin", userResponse.Role)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		c.SetParamValues("nonexistentuser")

		roleUpdateRequest := dto.UserRoleUpdateRequest{
			Role: "admin",
		}

		userResponse, err := userService.UpdateRoleUser(c, roleUpdateRequest)

		assert.Error(t, err)
		assert.Nil(t, userResponse)
		assert.Contains(t, err.Error(), "User not found")
	})
}
