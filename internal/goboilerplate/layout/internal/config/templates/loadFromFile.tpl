package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func loadConfigFromFile(cfg *Config) (err error) {
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; will read from env
			fmt.Println("[CONFIG] Config file not found.")
		} else {
			fmt.Printf("\n[CONFIG] Error while reading config file. Err: %s", err)
		}
	} else {
		if err = viper.Unmarshal(cfg); err != nil {
			fmt.Printf("\n[CONFIG] Error while unmarshalling config content. Err: %s", err)
		}
	}

	if err == nil {
		//WatchConfig, already reads new config, only need to re unmarshal to cfg.
		//Event triggers twice when using RichText Editors -- https://github.com/fsnotify/fsnotify/issues/324
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("[CONFIG] Config file reloaded.")
			err := viper.Unmarshal(cfg)
			if err != nil {
				fmt.Printf("\n[CONFIG] Error while unmarshalling config content. Err: %s", err)
			}
		})

		viper.WatchConfig()
	}

	return err
}