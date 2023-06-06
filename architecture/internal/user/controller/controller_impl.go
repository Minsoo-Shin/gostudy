package controller

import (
	"ggurugi/entity"
	"ggurugi/pkg/util"
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
	var request entity.UserSaveRequest

	if err := ctx.Bind(&request); err != nil {
		return err
	}

	if err := request.Valid(); err != nil {
		return err
	}

	if err := c.service.Save(ctx.Request().Context(), request); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// SignIn godoc
// @Summary      로그인
// @Description  로그인
// @Tags         user
// @Router /users/sign-in [post]
func (c controller) SignIn(ctx echo.Context) error {
	var request entity.UserSignInRequest

	if err := ctx.Bind(&request); err != nil {
		return err
	}

	if err := request.Valid(); err != nil {
		return err
	}

	userInfo, err := c.service.SignIn(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	util.InitEmptySlice(&userInfo)

	return ctx.JSON(http.StatusOK, userInfo)
}

// GetUser godoc
// @Summary      유저 조회 (학생, 학부모, 선생)
// @Description  유저 조회 (학생, 학부모, 선생)
// @Tags         user
// @Success 200
// @Router /users [get]
func (c controller) GetUser(ctx echo.Context) error {
	var request entity.GetUserRequest
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
		request entity.UpdateUserRequest
		err     error
	)
	context := c.ctxutil.NewContextFromEcho(ctx)

	if err = ctx.Bind(&request); err != nil {
		return err
	}

	request.ID, err = c.ctxutil.GetUserID(context)
	if err != nil {
		return err
	}

	role, err := c.ctxutil.GetRole(context)
	if err != nil {
		return err
	}

	if err = request.Valid(role); err != nil {
		return err
	}

	if err := c.service.Update(context, request); err != nil {
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
		request entity.UserDeleteRequest
		err     error
	)
	context := c.ctxutil.NewContextFromEcho(ctx)

	if err := ctx.Bind(&request); err != nil {
		return err
	}

	request.ID, err = c.ctxutil.GetUserID(context)
	if err != nil {
		return err
	}

	if err := request.Valid(); err != nil {
		return err
	}

	if err := c.service.Delete(ctx.Request().Context(), request); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c controller) CheckDuplicatedUserID(ctx echo.Context) error {
	var request entity.CheckDuplicatedUserIDRequest

	if err := ctx.Bind(&request); err != nil {
		return err
	}

	if err := request.Valid(); err != nil {
		return err
	}

	isDuplicated, err := c.service.CheckDuplicatedUserID(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, isDuplicated)
}

func (c controller) CheckDuplicatedEmail(ctx echo.Context) error {
	var request entity.CheckDuplicatedEmailRequest

	if err := ctx.Bind(&request); err != nil {
		return err
	}

	if err := request.Valid(); err != nil {
		return err
	}

	isDuplicated, err := c.service.CheckDuplicatedEmail(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, isDuplicated)
}

func (c controller) CheckExistingTeacherUserID(ctx echo.Context) error {
	var request entity.CheckExistingTeacherUserIDRequest

	if err := ctx.Bind(&request); err != nil {
		return err
	}

	if err := request.Valid(); err != nil {
		return err
	}

	exist, err := c.service.CheckExistingTeacherUserID(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, exist)
}
