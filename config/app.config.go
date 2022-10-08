package config

type AppConfig struct {
	AppPort string `env:"APP_PORT" envDefault:"9999"`
	AppEnv  string `env:"APP_ENV" envDefault:"local"`
	Mode    string `env:"MODE" envDefault:"debug"` //value = debug || test || release
}
