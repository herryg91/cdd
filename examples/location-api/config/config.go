package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServiceName string `envconfig:"service_name" default:"location-api"`
	Environment string `envconfig:"environment" default:"dev"`
	Maintenance bool   `envconfig:"maintenance" default:"false"`
	RestPort    int    `envconfig:"rest_port" default:"18080" required:"true"`
	GrpcPort    int    `envconfig:"grpc_port" default:"19090" required:"true"`

	DBHost         string `envconfig:"DB_HOST" default:"localhost"`
	DBPort         string `envconfig:"DB_PORT" default:"3306"`
	DBUserName     string `envconfig:"DB_USERNAME" default:"root"`
	DBPassword     string `envconfig:"DB_PASSWORD" default:"root"`
	DBDatabaseName string `envconfig:"DB_DBNAME" default:"cdd"`
	DBLogMode      bool   `envconfig:"DB_LOG_MODE" default:"true"`
}

func New() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)

	return cfg
}
