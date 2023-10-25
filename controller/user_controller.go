package controller

import (
	"fmt"
	"madyasantosa/ruangkegiatan/dto"
	"madyasantosa/ruangkegiatan/helper"
	"madyasantosa/ruangkegiatan/service"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	GetAllUsers(ctx echo.Context) error
	GetUserByUsername(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
	RegisterUser(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	UpdateRoleUser(ctx echo.Context) error
	ChangePassword(ctx echo.Context) error
	UserLogin(ctx echo.Context) error
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(us service.UserService) UserController {
	return &UserControllerImpl{UserService: us}
}

func (uc *UserControllerImpl) GetAllUsers(ctx echo.Context) error {
	params := ctx.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		fmt.Println(err)
		return helper.StatusBadRequest(ctx, fmt.Errorf("Params limit not valid"))
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		fmt.Println(err)
		return helper.StatusBadRequest(ctx, fmt.Errorf("Params offset not valid"))
	}
	res, total, err := uc.UserService.GetAllUsers(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Users not found") {
			return helper.StatusNotFound(ctx, err)
		}
		return helper.StatusInternalServerError(ctx, err)
	}

	return helper.StatusOKWithPagination(ctx, "Success to Get Data", res, offset, limit, total)
}

func (uc *UserControllerImpl) GetUserByUsername(ctx echo.Context) error {
	res, err := uc.UserService.GetUserByUsername(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return helper.StatusNotFound(ctx, fmt.Errorf("User %s Not Found!", ctx.Param("username")))
		}
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to Get User %s", ctx.Param("username")))
	}

	return helper.StatusOK(ctx, "Success to get data", res)
}
func (uc *UserControllerImpl) DeleteUser(ctx echo.Context) error {
	err := uc.UserService.DeleteUser(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "User not found") {
			return helper.StatusNotFound(ctx, fmt.Errorf("User %s not found!", ctx.Param("username")))
		}
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to delete user %s", ctx.Param("username")))
	}

	return helper.StatusNoContent(ctx, "Success to delete data")
}

func (uc *UserControllerImpl) RegisterUser(ctx echo.Context) error {
	userCreateRequest := dto.UserCreateRequest{}
	err := ctx.Bind(&userCreateRequest)
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}
	res, err := uc.UserService.CreateUser(ctx, userCreateRequest)

	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}
		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed to create user"))
	}

	return helper.StatusCreated(ctx, "Success to created user", res)
}

func (uc *UserControllerImpl) UpdateUser(ctx echo.Context) error {
	userUpdateRequest := dto.UserUpdateRequest{}
	err := ctx.Bind(&userUpdateRequest)
	fmt.Println(userUpdateRequest)
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}
	res, err := uc.UserService.UpdateUser(ctx, userUpdateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "User not found") {
			return helper.StatusNotFound(ctx, fmt.Errorf("User %s Not found!", ctx.Param("username")))
		}

		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed update user %s", ctx.Param("username")))
	}
	return helper.StatusOK(ctx, "Success to update user", res)

}
func (uc *UserControllerImpl) UpdateRoleUser(ctx echo.Context) error {
	userRoleUpdateRequest := dto.UserRoleUpdateRequest{}
	err := ctx.Bind(&userRoleUpdateRequest)
	fmt.Println(userRoleUpdateRequest)
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}
	res, err := uc.UserService.UpdateRoleUser(ctx, userRoleUpdateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "User not found") {
			return helper.StatusNotFound(ctx, fmt.Errorf("User %s not found!", ctx.Param("username")))
		}

		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed update role user %s", ctx.Param("username")))
	}
	return helper.StatusOK(ctx, "Success to update role user", res)

}
func (uc *UserControllerImpl) ChangePassword(ctx echo.Context) error {
	authorization := ctx.Request().Header["Authorization"]
	userToken := strings.Split(authorization[0], " ")[1]
	data, err := helper.ExtractToken(userToken)

	changePasswordRequest := dto.ChangePasswordRequest{}
	err = ctx.Bind(&changePasswordRequest)
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}
	err = uc.UserService.ChangePassword(ctx, changePasswordRequest, data.Username)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "User not found") {
			return helper.StatusNotFound(ctx, fmt.Errorf("User %s not found!", data.Username))
		}

		return helper.StatusInternalServerError(ctx, fmt.Errorf("Failed change password user %s", data.Username))
	}
	return helper.StatusNoContent(ctx, "Success to change password user")

}
func (uc *UserControllerImpl) UserLogin(ctx echo.Context) error {
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
