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

func TestUserServiceImpl_GetAllUsers(t *testing.T) {
	userService := &service.UserServiceImpl{
		UserRepository: &MockUsersRepository{
			GetAllUsersFunc: func(offset int, limit int, search string) ([]model.User, int, error) {
				if search == "test" {
					return []model.User{
						{Username: "user1"},
						{Username: "user2"},
					}, 2, nil
				} else if search == "empty" {
					return []model.User{}, 0, nil
				} else {
					return nil, 0, fmt.Errorf("Error: Users not found")
				}
			},
		},
	}

	e := echo.New()
	rec := httptest.NewRecorder()

	t.Run("SearchTermProvided", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/your-api-endpoint?s=test&limit=10&offset=0", nil)
		c := e.NewContext(req, rec)

		users, total, err := userService.GetAllUsers(c)

		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, 2, total)
		assert.Len(t, users, 2)
	})

	t.Run("EmptySearchTerm", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/your-api-endpoint?limit=10&offset=0&s=", nil)
		c := e.NewContext(req, rec)

		users, total, _ := userService.GetAllUsers(c)

		assert.Equal(t, 0, total)
		assert.Len(t, users, 0)
	})

	t.Run("ErrorInSearch", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/your-api-endpoint?s=error&limit=10&offset=0", nil)
		c := e.NewContext(req, rec)

		users, _, err := userService.GetAllUsers(c)

		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Contains(t, err.Error(), "Users not found")
	})
}
