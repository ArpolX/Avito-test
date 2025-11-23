package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServerConfig
	DB Postgres
}

type ServerConfig struct {
	ServerHost string `env:"SERVER_HOST,required"`
	ServerPort string `env:"SERVER_PORT,required"`
}

type Postgres struct {
	HostDB     string `env:"DB_HOST,required"`
	PortDB     string `env:"DB_PORT,required"`
	PasswordDB string `env:"DB_PASSWORD,required"`
	NameDB     string `env:"DB_NAME,required"`
	UserDB     string `env:"DB_USER,required"`
}

func NewConfig() (Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to read env: %w", err)
	}
	return cfg, nil
}
