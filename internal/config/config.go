package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type config struct {
	viper  *viper.Viper
	Values *baseValues
}

func (c *config) init() error {
	c.viper.SetConfigName("dotman.conf")
	c.viper.SetConfigType("toml")
	c.viper.AddConfigPath("$HOME/.config/dotman")
	c.viper.AddConfigPath("/etc/dotman")
	c.viper.AddConfigPath(".")

	if err := c.viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	values, err := newBaseValues(c.viper)
	if err != nil {
		return fmt.Errorf("error parsing config: %v", err)
	}
	c.Values = values

	return nil
}

func (c *config) String() string {
	return fmt.Sprintf("{Values: %s}", c.Values)
}

func newConfig() *config {
	return &config{
		viper: viper.New(),
	}
}

var (
	activeConfig *config = nil
	initError    error   = nil
	once         sync.Once
)

func Config() (*config, error) {
	once.Do(func() {
		activeConfig = newConfig()
		initError = activeConfig.init()
	})
	if initError != nil {
		return nil, initError
	}
	return activeConfig, nil
}
