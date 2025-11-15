package constants

const (
	// JWT
	JWT_SECRET_ENV     = "JWT_SECRET"
	DEFAULT_JWT_SECRET = "sagara-tech-secret-key-2025"
	JWT_EXPIRY_HOURS   = 24

	// Token types
	TOKEN_TYPE_BEARER = "Bearer"

	// Auth contexts
	CTX_USER_ID    = "userID"
	CTX_USER_EMAIL = "email"
	CTX_USER_ROLE  = "role"

	// Password
	MIN_PASSWORD_LENGTH = 6
	BCRYPT_COST         = 10
)
