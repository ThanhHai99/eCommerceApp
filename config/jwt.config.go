package config

type JwtConfigs struct {
	SecretKey string `env:"JWT_SECRET_KEY" envDefault:"secret-key"`
}
