package response

var (
	MsgInternalServerError = "服务器内部错误"
	MsgNoRoute             = "路径不存在"
	MsgNoMethod            = "非法请求方式"
	MsgForbidden           = "无访问权限"
	MsgUnauthorized        = "用户未登录"
	MsgInvalidParam        = "请求参数错误"

	MsgLoginSuccess         = "用户登录成功"
	MsgLogoutSuccess        = "用户注销成功"
	MsgNoLogoutToken        = "用户注销token不存在"
	MsgLoginFailure         = "用户名或密码错误"
	MsgLoginInvalidCaptcha  = "验证码错误"
	MsgLoginUsingBannedUser = "用户被禁用"
)
