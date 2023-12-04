package cons

const (
	// Postgre
	ConnectionStringTemplate = "host=%s user=%s password=%s dbname=%s port=%s timezone=Asia/Jakarta sslmode=disable"

	// Redis
	RedisConnectionTemplate      = "%s:%s"
	RedisRefreshTokenTemplate    = "refresh_token:%s"
	RedisVerifyEmailCodeTemplate = "verify_email_code:%s"

	// Mongo
	PhotoIDTempate = "%d:%s"
)
