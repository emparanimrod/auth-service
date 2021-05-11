package core

import (
	"fmt"

	"auth/configs"
)

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (d Database) String(sslmode string) string {
	return fmt.Sprintf(""+
		"host=%s "+
		"port=%s "+
		"user=%s "+
		"dbname=%s "+
		"password=%s "+
		"sslmode=%v"+
		"", d.Host, d.Port, d.User, d.DBName, d.Password, sslmode)
}

type Config struct {
	DB Database

	Secret        string
	TokenDuration uint

	GRPCPort string
}

func GetConfig(cfg configs.EnvVarConfig) *Config {
	return &Config{
		DB: Database{
			User:     cfg.DBUser,
			Password: cfg.DBPassword,
			Host:     cfg.DBHost,
			Port:     cfg.DBPort,
			DBName:   cfg.DBName,
		},
		Secret:        cfg.SecretKey,
		TokenDuration: cfg.TokenDuration,
		GRPCPort:      cfg.GRPCPort,
	}
}
