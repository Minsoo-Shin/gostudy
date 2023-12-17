package controller

import (
	"github.com/Minsoo-Shin/go-boilerplate/domain"
	"github.com/Minsoo-Shin/go-boilerplate/pkg/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateUser godoc
// @Summary      유저 생성 (학생, 학부모, 선생)
// @Description  유저 생성 (학생, 학부모, 선생)
// @Tags         user
// @Success 204
// @Router /users [post]
func (c controller) CreateUser(ctx echo.Context) error {
	var request domain.UserCreateRequest

	if err := ctx.Bind(&request); err != nil {
		return err
	}

	if err := c.service.Create(ctx.Request().Context(), request); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// GetUser godoc
// @Summary      유저 조회 (학생, 학부모, 선생)
// @Description  유저 조회 (학생, 학부모, 선생)
// @Tags         user
// @Success 200
// @Router /users [get]
func (c controller) GetUser(ctx echo.Context) error {
	var request domain.UserFindRequest
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	if err := request.Valid(); err != nil {
		return err
	}

	userInfo, err := c.service.Find(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	util.InitEmptySlice(&userInfo)

	return ctx.JSON(http.StatusOK, userInfo)
}

// UpdateUser godoc
// @Summary      유저 정보 업데이트
// @Description  Parameter 참조
// @Tags         user
// @Success 204
// @Router /users/{userID} [put]
func (c controller) UpdateUser(ctx echo.Context) error {
	var (
		request domain.UserUpdateRequest
		err     error
	)
	if err = ctx.Bind(&request); err != nil {
		return err
	}

	if err := c.service.Update(ctx.Request().Context(), request); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// DeleteUser godoc
// @Summary      유저 탈퇴
// @Description  Parameter 참조
// @Tags         user
// @Success 204
// @Router /users [delete]
func (c controller) DeleteUser(ctx echo.Context) error {
	var (
		request domain.UserDeleteRequest
	)

	if err := ctx.Bind(&request); err != nil {
		return err
	}

	if err := c.service.Delete(ctx.Request().Context(), request); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
