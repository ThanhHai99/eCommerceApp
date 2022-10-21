package config

type MailConfig struct {
	SmtpHost          string `env:"SMTP_HOST" envDefault:"smtp.gmail.com"`
	SmtpPort          int    `env:"SMTP_PORT" envDefault:"587"`
	SmtpUser          string `env:"SMTP_USER" envDefault:"bipbip.shopping@gmail.com"`
	SmtpPass          string `env:"SMTP_PASS" envDefault:"tnvzpstmntmakndr"`
	SmtpDisplaySender string `env:"SMTP_DISPLAY_SENDER" envDefault:"Hehe boi"`
}
