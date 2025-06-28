package errno

var (
	Success              = NewErrNo(SuccessCode, SuccessMsg)
	InternalServiceError = NewErrNo(InternalServiceErrorCode, "internal server error")

	AuthInvalid        = NewErrNo(AuthInvalidCode, "authentication failure")
	AuthAccessExpired  = NewErrNo(AuthAccessExpiredCode, "token expiration")
	AuthRefreshExpired = NewErrNo(AuthRefreshExpiredCode, "token refresh expired")
	AuthNoToken        = NewErrNo(AuthNoTokenCode, "lack of token")
)
