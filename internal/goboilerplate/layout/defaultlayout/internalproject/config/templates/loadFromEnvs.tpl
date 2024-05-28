package config

import "github.com/kelseyhightower/envconfig"

const envPrefix string = "{{.}}"

func loadConfigFromEnvs(cfg *Config) (err error) {
	return envconfig.Process(envPrefix, cfg)
}
