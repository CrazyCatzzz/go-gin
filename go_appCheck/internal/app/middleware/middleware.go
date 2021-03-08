package middleware

import (
	"appcheck/internal/app/response"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool

// NoRoute 路径不存在
func NoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.NewErrorResponse(response.MsgNoRoute))
		c.Abort()
	}
}

// NoMethod 方法不存在
func NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(response.MsgNoRoute))
		c.Abort()
	}
}

// AllowPathPrefixSkipper 请求路径 白名单
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		urlPath := c.Request.URL.Path
		pathLen := len(urlPath)
		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && urlPath[:pl] == p {
				return true
			}
		}
		return false
	}
}

// AllowPathPrefixNoSkipper 请求路径 黑名单
func AllowPathPrefixNoSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		uPath := c.Request.URL.Path
		pathLen := len(uPath)
		for _, prefix := range prefixes {
			if prefixLen := len(prefix); pathLen >= prefixLen && uPath[:prefixLen] == prefix {
				return false
			}
		}
		return true
	}
}

// AllowMethodAndPathPrefixSkipper 请求方法+请求路径 白名单
func AllowMethodAndPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		urlPath := JoinRouter(c.Request.Method, c.Request.URL.Path)
		pathLen := len(urlPath)

		for _, prefix := range prefixes {
			if prefixLen := len(prefix); pathLen >= prefixLen && urlPath[:prefixLen] == prefix {
				return true
			}
		}
		return false
	}
}

// AllowRemoteAddressSkipper IP白名单
func AllowRemoteAddressSkipper(address ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		remoteAddr := c.ClientIP()
		for _, addr := range address {
			if addr == remoteAddr {
				return true
			}
		}
		return false
	}
}

// AllowRemoteAddressNoSkipper IP黑名单
func AllowRemoteAddressNoSkipper(address ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		remoteAddr := c.ClientIP()
		for _, addr := range address {
			if addr == remoteAddr {
				return false
			}
		}
		return true
	}
}

// JoinRouter 方法与路径相连
func JoinRouter(m, p string) string {
	if len(p) > 0 && p[0] != '/' {
		p = "/" + p
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(m), path.Clean(p))
}

// SkipHandler 统一处理跳过函数
func SkipHandler(c *gin.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(c) {
			return true
		}
	}

	return false
}
