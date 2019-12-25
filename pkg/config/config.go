package config

import (
	"fmt"
	"os"

	chassisconfig "pgxs.io/chassis/config"
)

const (
	configFileEnvKey = "PG_CONF_FILE"
)

func init() {
	//todo check config
	loadFromEnv()
}

//CustomConfig all config
type CustomConfig struct {
	chassisconfig.Config `yaml:",inline"`
	Pkg                  PkgConfig
}

//PkgConfig LDAP config
type PkgConfig struct {
	Host string
}

var config CustomConfig

//Pkg gopkg settings
func Pkg() *PkgConfig {
	return &config.Pkg
}

//loadFromEnv load config file from env
func loadFromEnv() {
	fileName := os.Getenv(configFileEnvKey)
	if err := chassisconfig.LoadConfigFromFile(fileName, &config); err != nil {
		fmt.Printf("load file config error: %s\n", err)
		os.Exit(1)
	}
}
