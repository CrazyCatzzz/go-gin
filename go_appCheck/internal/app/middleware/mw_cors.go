package middleware

import (
	"github.com/gin-gonic/gin"
)

// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
