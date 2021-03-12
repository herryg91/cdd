package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServiceName string `envconfig:"service_name" default:"users-api"`
	Environment string `envconfig:"environment" default:"dev"`
	Maintenance bool   `envconfig:"maintenance" default:"false"`
	RestPort    int    `envconfig:"rest_port" default:"18081" required:"true"`
	GrpcPort    int    `envconfig:"grpc_port" default:"19091" required:"true"`

	DBHost         string `envconfig:"DB_HOST" default:"localhost"`
	DBPort         string `envconfig:"DB_PORT" default:"3306"`
	DBUserName     string `envconfig:"DB_USERNAME" default:"root"`
	DBPass         string `envconfig:"DB_PASS" default:"root"`
	DBDatabaseName string `envconfig:"DB_DBNAME" default:"cdd"`
	DBLogMode      bool   `envconfig:"DB_LOG_MODE" default:"true"`

	LocationApi string `envconfig:"LOCATION_API" default:"localhost:19090"`
}

func New() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)

	return cfg
}
