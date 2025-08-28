package configs

type PostgresConfig struct {
	UserDB        string
	PasswordDB    string
	NameDB        string
	HostDB        string
	PortDB        string
	SSLModeDB     string
	UserAppDB     string
	PasswordAppDB string
	NameAppDB     string
}

func NewPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		UserDB:        getEnv("DB_USER", "postgres"),
		PasswordDB:    getEnv("DB_PASSWORD", "supersecret"),
		NameDB:        getEnv("DB_NAME", "postgres"),
		HostDB:        getEnv("DB_HOST", "localhost"),
		PortDB:        getEnv("DB_PORT", "5432"),
		SSLModeDB:     getEnv("DB_SSLMODE", "disable"),
		UserAppDB:     getEnv("APP_DB_USER", "app"),
		PasswordAppDB: getEnv("APP_DB_PASSWORD", "secret"),
		NameAppDB:     getEnv("APP_DB_NAME", "wb_level0_db"),
	}
}
