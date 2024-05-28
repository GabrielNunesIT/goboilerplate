package config

import (
	"fmt"
	"sync"
	"time"
	"fmt" 

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {}

var cfg *Config = &Config{}

var onceInit sync.Once

func initConfig() {
	viper.SetConfigName("{{.}}")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	loadConfigFromFile(cfg)
	loadConfigFromEnvs(cfg)

	fmt.Println("[CONFIG] Config loaded successfully.")
}

func GetConfig() *Config {
	onceInit.Do(func() {
		initConfig()
	})
	return cfg
}