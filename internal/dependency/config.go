package dependency

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		App        app
		Rest       rest
		PostgreDB  postgreDB
		MongoDB    mongoDB
		RedisCache redisconfig
		Jwt        jwt
		Email      email
	}

	app struct {
		AppName         string `env:"APP_NAME"`
		GracefulTimeout uint   `env:"GRACEFUL_TIMEOUT"`
		Domain          string `env:"BASE_URL"`
		EnvMode         string `env:"ENV_MODE"`
	}

	rest struct {
		RequestTimeout uint `env:"REQUEST_TIMEOUT"`
		Port           uint `env:"REST_PORT"`
	}

	postgreDB struct {
		DBHost string `env:"DB_HOST"`
		DBPort string `env:"DB_PORT"`
		DBName string `env:"DB_NAME"`
		DBUser string `env:"DB_USER"`
		DBPass string `env:"DB_PASS"`
		DBTz   string `env:"DB_TZ" env-default:"Asia/Jakarta"`
	}

	mongoDB struct {
		MongoUri string `env:"MONGODB_URI"`
	}

	redisconfig struct {
		HOST     string `env:"REDIS_HOST"`
		PORT     string `env:"REDIS_PORT"`
		Password string `env:"REDIS_PASSWORD"`
	}

	jwt struct {
		JWTSecret              string `env:"JWT_SECRET"`
		AccessTokenExpiration  uint   `env:"ACCESS_TOKEN_EXPIRATION"`
		RefreshTokenExpiration uint   `env:"REFRESH_TOKEN_EXPIRATION"`
	}

	email struct {
		SenderName     string `env:"EMAIL_SENDER_NAME"`
		SenderAddress  string `env:"EMAIL_SENDER_ADDRESS"`
		SenderPassword string `env:"EMAIL_SENDER_PASSWORD"`
	}
)

func NewConfig(logger Logger) (*Config, error) {
	config := new(Config)

	err := cleanenv.ReadEnv(config)
	if err != nil {
		logger.Fatalf("Failed to load config")
		return nil, err
	}

	logger.Infof("Successfully load config", nil)

	return config, err
}
