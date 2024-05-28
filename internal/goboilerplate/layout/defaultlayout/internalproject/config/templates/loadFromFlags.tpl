package config

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func addFlags() {
	// Add Config flags here...
	// Need to match field tag "mapstructure" to use auto unmarshall | Use "." for nested structs. Ex: webServer.port
}

func loadConfigFromFlag(cfg *Config) (err error) {
	addFlags()

	flag.VisitAll(func (f *flag.Flag) {
		if f.value.String() == f.DefValue {
			pflag.CommandLine.AddGoFlag(f)
		}
	})

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("\n[CONFIG] Error while reading config flags. Err: %s", err)
	} else {
		if err = viper.Unmarshal(cfg); err != nil {
			fmt.Printf("\n[CONFIG] [CONFIG] Error while unmarshalling config content. Err: %s", err)
		}
	}

	return err
}