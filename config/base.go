package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type baseConfig struct {
	viper  *viper.Viper
	Values *baseConfigValues
}

func (c *baseConfig) init() error {
	c.viper.SetConfigName("dotman.conf")
	c.viper.SetConfigType("toml")
	c.viper.AddConfigPath("$HOME/.config/dotman")
	c.viper.AddConfigPath("/etc/dotman")
	c.viper.AddConfigPath(".")

	if err := c.viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	values, err := newBaseConfigValues(c.viper)
	if err != nil {
		return fmt.Errorf("error parsing config: %v", err)
	}
	c.Values = values

	return nil
}

func newBaseConfig() *baseConfig {
	viper := viper.New()

	return &baseConfig{
		viper: viper,
	}
}

var config *baseConfig = nil
var initError error = nil
var once sync.Once

func BaseConfig() (*baseConfig, error) {
	once.Do(func() {
		config = newBaseConfig()
		initError = config.init()
	})
	if initError != nil {
		return nil, initError
	}
	return config, nil
}
