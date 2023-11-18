package constant

const (
	ConnectionStringTemplate = "host=%s user=%s password=%s dbname=%s port=%s timezone=Asia/Jakarta sslmode=disable"
	RedisConnectionTemplate  = "%s:%s"

	RedisRefreshTokenTemplate = "refresh_token:%s"
	RedisPaymentTokenTemplate = "payment_token:%s"
)
