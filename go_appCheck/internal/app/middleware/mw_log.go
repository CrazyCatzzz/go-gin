package middleware

/*
	日志中间件,支持按日期自动分割
*/

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	accessName   = "access.log"
	maxAge       = 3600 * 24 * 90
	rotationTime = 3600 * 24
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	writer, _ := rotatelogs.New(
		accessName+".%Y-%m-%d",
		rotatelogs.WithLinkName(accessName),
		rotatelogs.WithMaxAge(maxAge*time.Second),
		rotatelogs.WithRotationTime(rotationTime*time.Second),
	)
	logger.Out = writer
}

// Logger 日志中间件
func Logger(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		// 开始时间
		start := time.Now()

		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		// 参数记录
		latency := end.Sub(start)
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		logger.Infof("| %3d | %10v | %10s | %-4s | %s",
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
