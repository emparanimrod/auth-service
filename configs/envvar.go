package configs

import "github.com/kelseyhightower/envconfig"

type EnvVarConfig struct {
	// Database connection env configs
	DBUser     string `required:"true"`
	DBPassword string `required:"true"`
	DBHost     string `required:"true"`
	DBPort     string `required:"true"`
	DBName     string `required:"true"`

	// application specific configs
	SecretKey     string `required:"true"`
	TokenDuration uint   `required:"true" default:"2880"`

	// application grpc port
	GRPCPort string `envconfig:"grpc_port" default:"7301"`
}

func GetEnvConfig() (*EnvVarConfig, error) {
	var cfg EnvVarConfig
	err := envconfig.Process("authapp", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
