package router

import (
	_ "appcheck/docs"

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
}

func (a *Router) Register(engine *gin.Engine) {
	a.registerAPI(engine)
}

func (a *Router) Prefix() string {
	return "/api/v1"
}
