package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

const (
	DbDriver = "mysql"
)

type Config struct {
	Env  string `env:"ENV,default=development"`
	Port string `env:"PORT,default=8080"`
	DB   Database
	Jwt  Jwt
}

type Jwt struct {
	Secret string `env:"JWT_SECRET"`
}

type Database struct {
	Host     string `env:"DATABASE_HOST,default=localhost"`
	Port     string `env:"DATABASE_PORT,default=5432"`
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Name     string `env:"DATABASE_NAME,required"`
}

func New(envFile string) (Config, error) {
	var cfg Config
	if err := godotenv.Load(envFile); err != nil {
		return cfg, err
	}

	if err := envdecode.Decode(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
