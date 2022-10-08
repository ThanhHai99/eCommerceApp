package config

type DatabaseConfig struct {
	DBType   string `env:"DB_TYPE" envDefault:"postgres"`
	DBHost   string `env:"DB_HOST" envDefault:"localhost"`
	DBPort   string `env:"DB_PORT" envDefault:"5432"`
	DBUser   string `env:"DB_USER" envDefault:"postgres"`
	DBPass   string `env:"DB_PASS" envDefault:"root"`
	DBName   string `env:"DB_NAME" envDefault:"eCommerce"`
	DBSchema string `env:"DB_SCHEMA" envDefault:"public"`
	SSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
}
