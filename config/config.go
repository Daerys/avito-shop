package config

import "github.com/caarlos0/env/v8"

// Config содержит настройки приложения.
type Config struct {
	GRPC GRPC
	PG   PG
	JWT  JWT
}

// GRPC содержит настройки для gRPC и gRPC Gateway.
type GRPC struct {
	Port        string `env:"GRPC_PORT" envDefault:"50051"`
	GatewayPort string `env:"GRPC_GATEWAY_PORT" envDefault:"8080"`
}

// PG содержит настройки подключения к PostgreSQL.
type PG struct {
	URL      string
	Host     string `env:"DATABASE_HOST" envDefault:"localhost"`
	Port     string `env:"DATABASE_PORT" envDefault:"5432"`
	DB       string `env:"DATABASE_NAME" envDefault:"shop"`
	User     string `env:"DATABASE_USER" envDefault:"postgres"`
	Password string `env:"DATABASE_PASSWORD" envDefault:"password"`
}

// JWT содержит настройки для JWT аутентификации.
type JWT struct {
	Secret string `env:"JWT_SECRET" envDefault:"f1ae64498f9910d0e94a22c1e9210c8432831e53d47cbcc784b59eb58a6a88f2caf4cfe815802d3a541154aec3196b8ed809db84890755bfc8e77bf8b0243216908a296628970dcfdd454d337a03eedb9437a09ba511ce146d21c2b024fb2b9ea2293ad79c687f65df9eb16c63e595ab8cbbca9c1dcb533ffd356fc974b66602"`
}

// NewConfig загружает конфигурацию из переменных окружения.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
