package consts

const (
	PageSize         = 10
	PhoneVerifyType1 = 1            // 1-注册/登录
	PhoneVerifyType2 = 2            // 2-找回密码
	ContextKey       = "ContextKey" // 上下文变量存储键名，前后端系统共享

	// for 登录相关
	TokenType          = "Bearer"
	CacheModeRedis     = 2
	GtokenWriterPrefix = "Writer:" // gtoken登录，小说作者后台前缀区分
	BackendServerName  = "小说系统"
	MultiLogin         = true
	FrontendMultiLogin = true
	GTokenExpireIn     = 10 * 24 * 60 * 60

	// for writer
	CtxWriterId      = "CtxAdminId"
	CtxWriterName    = "CtxAdminName"
	CtxWriterIsAdmin = "CtxAdminIsAdmin"
	CtxWriterRoleIds = "CtxAdminRoleIds"
)
