package errno

const (
	SuccessCode = 10000
	SuccessMsg  = "success"
)

const (
	ParamVerifyErrorCode  = 20000 + iota
	ParamMissingErrorCode //参数缺失
)

const (
	AuthInvalidCode        = 30000 + iota // 鉴权失败
	AuthAccessExpiredCode                 // 访问令牌过期
	AuthRefreshExpiredCode                // 刷新令牌过期
	AuthNoTokenCode                       // 没有 token
)

const (
	InternalServiceErrorCode = 50000 + iota // 内部服务错误
	InternalDatabaseErrorCode
)
