package ginplus

import (
	"github.com/gin-gonic/gin"
)

const (
	SessionUserID    = "X-USER-ID"
	SessionCaptchaID = "X-CAPTCHA-ID"
	CookieName       = "JSESSIONID"
)

func GetUserID(c *gin.Context) string {
	return c.GetString(SessionUserID)
}
