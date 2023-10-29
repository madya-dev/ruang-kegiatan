package test

import (
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/features/users/service"
	"madyasantosa/ruangkegiatan/model"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceImpl_CreateUser(t *testing.T) {
	t.Run("UserCreatedSuccessfully", func(t *testing.T) {
		userService := &service.UserServiceImpl{
			UserRepository: &MockUsersRepository{
				GetUserByUsernameFunc: func(username string) (*model.User, error) {
					return nil, nil
				},
				CreateUserFunc: func(user *model.User) (*model.User, error) {
					return user, nil
				},
			},
			Validate: validator.New(),
		}

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/your-api-endpoint", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		createUserRequest := dto.UserCreateRequest{
			Username:     "newuser",
			Password:     "newpassword123",
			FullName:     "John Doe",
			StudyProgram: "Computer Science",
			Phone:        "1234567890",
		}

		userResponse, err := userService.CreateUser(c, createUserRequest)

		assert.NoError(t, err)
		assert.NotNil(t, userResponse)
		assert.Equal(t, "newuser", userResponse.Username)
	})

	t.Run("UserAlreadyExists", func(t *testing.T) {
		userService := &service.UserServiceImpl{
			UserRepository: &MockUsersRepository{
				GetUserByUsernameFunc: func(username string) (*model.User, error) {
					return &model.User{}, nil
				},
			},
			Validate: validator.New(),
		}

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/your-api-endpoint", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		createUserRequest := dto.UserCreateRequest{
			Username:     "existinguser",
			Password:     "newpassword123",
			FullName:     "John Doe",
			StudyProgram: "Computer Science",
			Phone:        "1234567890",
		}

		userResponse, err := userService.CreateUser(c, createUserRequest)

		assert.Error(t, err)
		assert.Nil(t, userResponse)
		assert.Contains(t, err.Error(), "User already exists")
	})
}
