package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	defaultConfigType = "yml"
	defaultHostDB     = "localhost"
	defaultPortDB     = 3306
	defaultUserDB     = "api"
	defaultPasswordDB = "secret"
	defaultNameDB     = "paxful"
	defaultPortAPI    = 1357
	prefixEnvironmet  = "API"
)

// Config is the main config.
type Config struct {
	DB  DBConfig
	API APIConfig
}

type (
	DBConfig struct {
		Host     string `mapstructure:"HOST"`
		Port     int    `mapstructure:"PORT"`
		User     string `mapstructure:"USER"`
		Password string `mapstructure:"PASSWORD"`
		Name     string `mapstructure:"NAME"`
	}

	APIConfig struct {
		Host string `mapstructure:"HOST"`
		Port int    `mapstructure:"PORT"`
	}
)

// Init is uses as constructor for config.
func Init(filePath string) (conf *Config, err error) {
	dir, fileName, err := parseFilePath(filePath)
	if err != nil {
		return nil, err
	}

	var runtimeViper = viper.New()

	// Set default values and init settings.
	populateDefaults(runtimeViper)

	// Load from file.
	runtimeViper.AddConfigPath(dir)
	runtimeViper.SetConfigName(fileName)
	runtimeViper.SetConfigType(defaultConfigType)

	// Load from evironment if exists.
	runtimeViper.SetEnvPrefix(prefixEnvironmet)
	runtimeViper.AutomaticEnv()

	// Fill viper map.
	if err = runtimeViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return
		}
	}

	// Fill config from viper.
	err = runtimeViper.Unmarshal(&conf)
	return
}

// populateDefaults is used for setting the default settings.
func populateDefaults(viperRuntime *viper.Viper) {
	viperRuntime.SetDefault("DB.Host", defaultHostDB)
	viperRuntime.SetDefault("DB.Port", defaultPortDB)
	viperRuntime.SetDefault("DB.User", defaultUserDB)
	viperRuntime.SetDefault("DB.Password", defaultPasswordDB)
	viperRuntime.SetDefault("DB.Name", defaultNameDB)
	viperRuntime.SetDefault("API.Port", defaultPortAPI)
}

// parseFilePath is used for path conversion.
func parseFilePath(filePath string) (dir string, fileName string, err error) {
	path := strings.Split(filePath, "/")
	if len(path) < 2 {
		return "", "", ErrFileNotFound
	}
	dir = path[0]
	fileName = path[1]

	return
}
