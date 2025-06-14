package config

import (
	"fmt"

	"dotman/config/value"

	"github.com/spf13/viper"
)

type baseConfigValues struct {
	Giturl value.Value[string]
}

func (b *baseConfigValues) validate() error {
	if !b.Giturl.IsValid() {
		return fmt.Errorf("giturl is required")
	}
	return nil
}

func (b *baseConfigValues) String() string {
	return fmt.Sprintf("{Giturl: %s}", b.Giturl)
}

func newBaseConfigValues(viper *viper.Viper) (*baseConfigValues, error) {
	values := &baseConfigValues{
		Giturl: value.NewStringValue("giturl", "", true, viper),
	}
	if err := values.validate(); err != nil {
		return nil, fmt.Errorf("error validating config: %v", err)
	}
	return values, nil
}
