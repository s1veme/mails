package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDevelopment bool `env:"IS_DEVELOPMENT" env-default:"false"`
	HTTP          HTTPConfig
	Postgres      PostgresConfig
	SMTP          SMTPConfig
	Redis         RedisConfig
}

type HTTPConfig struct {
	Host           string   `env:"HTTP_HOST" env-default:"localhost"`
	Port           int      `env:"HTTP_PORT" env-default:"8080"`
	AllowedOrigins []string `env:"HTTP_ALLOWED_ORIGINS" env-default:"http://localhost:3000"`
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     int    `env:"POSTGRES_PORT" env-default:"5432"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DATABASE"`
}

type SMTPConfig struct {
	Host string `env:"SMTP_HOST"`
	Port int    `env:"SMTP_PORT"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST" env-default:"localhost"`
	Port     int    `env:"REDIS_PORT" env-default:"6379"`
	Password string `env:"REDIS_PASSWORD" env-default:""`
	Database int    `env:"REDIS_DB" env-default:"0"`
}

var instance *Config

const envDefaultPath = "./deployment/.env"

func GetConfig() *Config {
	instance = new(Config)

	if err := cleanenv.ReadConfig(envDefaultPath, instance); err != nil {
		helpText := "Error read env"
		help, _ := cleanenv.GetDescription(instance, &helpText)
		log.Print(help)
		log.Fatal(err)
	}

	return instance
}

func GetDevelopmentConfig() *Config {
	instance = new(Config)

	if err := cleanenv.ReadEnv(nil); err != nil {
		helpText := "Error read env"
		help, _ := cleanenv.GetDescription(instance, &helpText)
		log.Print(help)
		log.Fatal(err)
	}

	return instance
}
