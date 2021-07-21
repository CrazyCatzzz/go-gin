package router

import (
	"yh-server/docs"
	"yh-server/internal/app/api"
	"yh-server/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var _ IRouter = (*Router)(nil)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	// 路由注册
	Register(*gin.Engine)
	// 路由前缀
	Prefix() string
}

type Router struct {
	HomeAPI *api.HomeController
}

func (a *Router) Register(engine *gin.Engine) {
	docs.SwaggerInfo.Host = config.GetSwaggerHost()
	a.registerAPI(engine)
}

func (a *Router) Prefix() string {
	return "/api/v1"
}
