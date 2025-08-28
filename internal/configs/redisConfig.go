package configs

type RedisConfig struct {
	RedisAddr string
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		RedisAddr: getEnv("REDIS_ADDR", "localhost:6379"),
	}
}
