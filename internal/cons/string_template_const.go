package cons

const (
	// Postgre
	ConnectionStringTemplate = "host=%s user=%s password=%s dbname=%s port=%s timezone=Asia/Jakarta sslmode=disable"
	
	// Redis
	RedisConnectionTemplate  = "%s:%s"
	RedisRefreshTokenTemplate = "refresh_token:%s"
	RedisPaymentTokenTemplate = "payment_token:%s"

	// Mongo
	PhotoIDTempate = "%d:%s"
)
