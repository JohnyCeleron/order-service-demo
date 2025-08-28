package configs

type EnvironmentConfig struct {
	Environment string
}

func NewEnvironmentConfig() *EnvironmentConfig {
	return &EnvironmentConfig{
		Environment: getEnv("ENVIRONMENT", "local"),
	}
}
