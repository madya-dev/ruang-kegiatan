package service

import "github.com/labstack/echo/v4"

func (s *UserServiceImpl) DeleteUser(ctx echo.Context) error {
	err := s.UserRepository.DeleteUser(ctx.Param("username"))
	if err != nil {
		return err
	}

	return nil
}
