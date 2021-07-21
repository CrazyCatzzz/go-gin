package api

import (
	"net/http"
	"yh-server/internal/app/response"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var HomeControllerSet = wire.NewSet(wire.Struct(new(HomeController), "*"))

type HomeController struct {
}

// Login
// @Summary 登录
// @Description 登录
// @tags 登录
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response "{"success":true,"msg":"","data":null}"
// @Failure 400 {object} response.Response "{"success":false,"msg":"请求参数错误","data":null}"
// @Failure 401 {object} response.Response "{"success":false,"msg":"用户未登录","data":null}"
// @Failure 403 {object} response.Response "{"success":false,"msg":"访问受限","data":null}"
// @Failure 404 {object} response.Response "{"success":false,"msg":"路径不存在","data":null}"
// @Failure 500 {object} response.Response "{"success":false,"msg":"Internal Server Error","data":null}"
// @Router /login [post]
func (h *HomeController) Login(c *gin.Context) {

	c.JSON(http.StatusOK, response.NewSuccessResponseWithPayload(nil))
}
