package test

import (
	"fmt"
	"madyasantosa/ruangkegiatan/features/users/service"
	"madyasantosa/ruangkegiatan/model"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceImpl_GetUserByUsername(t *testing.T) {
	e := echo.New()

	mockRepo := &MockUsersRepository{}
	expectedUsername := "testuser"

	mockRepo.GetUserByUsernameFunc = func(username string) (*model.User, error) {
		if username == expectedUsername {
			return &model.User{
				Username: expectedUsername,
			}, nil
		}
		return nil, fmt.Errorf("User not found")
	}

	userService := &service.UserServiceImpl{
		UserRepository: mockRepo,
	}

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("username")
	c.SetParamValues(expectedUsername)

	t.Run("ValidUser", func(t *testing.T) {
		userResponse, err := userService.GetUserByUsername(c)

		assert.NoError(t, err)
		assert.NotNil(t, userResponse)
		assert.Equal(t, expectedUsername, userResponse.Username)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		invalidUsername := "nonexistentuser"
		c.SetParamValues(invalidUsername)

		userResponse, err := userService.GetUserByUsername(c)

		assert.Error(t, err)
		assert.Nil(t, userResponse)
		assert.Contains(t, err.Error(), "User not found")
	})
}
