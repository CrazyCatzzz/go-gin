package config

import "fmt"

type Swagger struct {
	Enabled bool   `toml:"enabled"`
	Host    string `toml:"host,omitempty"`
	Port    int    `toml:"port,omitempty"`
}

func GetSwaggerHost() string {
	if globalConfig.Swagger.Host == "" {
		globalConfig.Swagger.Host = globalConfig.Http.Host
	}

	if globalConfig.Swagger.Port == 0 {
		globalConfig.Swagger.Port = globalConfig.Http.Port
	}

	return fmt.Sprintf("%s:%d", globalConfig.Swagger.Host, globalConfig.Swagger.Port)
}

func EnableSwagger() bool {
	return globalConfig.Swagger.Enabled
}
