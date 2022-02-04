package config

import (
	"flag"

	"github.com/minipkg/log"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	LangEng = "eng"
	LangRus = "rus"
)

// Configuration is the struct for app configuration
type Configuration struct {
	Server struct {
		SMTPListen string
	}
	Log           log.Config
	EmailDefaults EmailDefaults
	APIProviders  []APIProvider
}

type EmailDefaults struct {
	SenderName  string
	SenderEmail string
}

type APIProvider struct {
	SysName  string
	HostName string
	APIKey   string
}

// defaultPathToConfig is the default path to the app config
const defaultPathToConfig = "config/config.yaml"

// Get func return the app config
func Get() (*Configuration, error) {
	// config is the app config
	var config Configuration = Configuration{}
	// pathToConfig is a path to the app config
	var pathToConfig string

	viper.AutomaticEnv() // read in environment variables that match
	//viper.BindEnv("pathToConfig")
	defPathToConfig := defaultPathToConfig
	if viper.Get("pathToConfig") != nil {
		defPathToConfig = viper.Get("pathToConfig").(string)
	}

	flag.StringVar(&pathToConfig, "config", defPathToConfig, "path to YAML/JSON config file")
	flag.Parse()

	if err := config.readConfig(pathToConfig); err != nil {
		return &config, err
	}

	return &config, nil
}

func (c *Configuration) readConfig(pathToConfig string) error {
	viper.SetConfigFile(pathToConfig)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.Errorf("Config file not found in %q", pathToConfig)
		} else {
			return errors.Errorf("Config file was found in %q, but was produced error: %v", pathToConfig, err)
		}
	}

	err := viper.Unmarshal(c)
	if err != nil {
		return errors.Errorf("Config unmarshal error: %v", err)
	}
	return nil
}
