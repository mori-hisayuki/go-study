package user

import "github.com/gin-gonic/gin"

type UserHandler struct{}

// GetUsers godoc
// @Summary ユーザー一覧取得
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} object "OK"
// @Router /v1/users [get]
func (h *UserHandler) GetUser(ctx *gin.Context) {}

// GetUserById godoc
// @Summary ユーザーの詳細情報を取得
// @Tags users
// @Accept json
// @Produce json
// @Param request path string true "ユーザーID"
// @Success 200 {object} object "OK"
// @Router /v1/users/:id [get]
func (h *UserHandler) GetUserById(ctx *gin.Context) {}

// EditUser godoc
// @Summary ユーザー情報の編集
// @Tags users
// @Accept json
// @Produce json
// @Param request body string true "ユーザーID"
// @Success 200 {object} object "OK"
// @Router /v1/users [post]
func (h *UserHandler) EditUser(ctx *gin.Context) {}

// DeleteUser godoc
// @Summary ユーザー情報の削除
// @Tags users
// @Accept json
// @Produce json
// @Param request path string true "ユーザーID"
// @Success 200 {object} object "OK"
// @Router /v1/users/:id [delete]
func (h *UserHandler) DeleteUser(ctx *gin.Context) {}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
