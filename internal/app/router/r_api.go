package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func (a *Router) registerAPI(engine *gin.Engine) {

	// var skippers []string
	// skippers = append(skippers, a.Prefix()+"/pub")

	// 资源校验中间件
	// engine.Use(middleware.Casbin(a.CasbinEnforcer, a.Prefix(), middleware.AllowPathPrefixSkipper(skippers...)))

	// 身份认证中间件
	store := cookie.NewStore([]byte("xaware"))
	engine.Use(sessions.Sessions("JSESSIONID", store))
	// engine.Use(middleware.Auth(a.AuthModel, middleware.AllowPathPrefixSkipper(skippers...)))

	// group := engine.Group(a.Prefix())
	a.registerHome(&engine.RouterGroup)

	// a.registerLogin(group)
	// a.RegisterLog(group)
}
