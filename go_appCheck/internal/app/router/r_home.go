package router

import "github.com/gin-gonic/gin"

func (a *Router) registerHome(engine *gin.RouterGroup) {
	engine.POST("/login", a.HomeAPI.Login)
}
