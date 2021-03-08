package injector

import (
	"appcheck/internal/app/middleware"
	"appcheck/internal/app/router"
	"appcheck/internal/config"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化Http Handle
func InitWebEngine(router router.IRouter) *gin.Engine {

	// debug
	if config.EnableHttpDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	// 404中间件
	engine.NoRoute(middleware.NoRoute())
	// 非法请求中间件
	engine.NoMethod(middleware.NoMethod())
	// 异常中间件
	engine.Use(gin.Recovery())
	// 日志中间件
	engine.Use(middleware.Logger())

	// swagger文档
	if config.EnableHttpDebug() {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	router.Register(engine)

	return engine
}
