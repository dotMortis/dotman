package config

import (
	"fmt"

	"dotman/internal/config/value"

	"github.com/spf13/viper"
)

type baseValues struct {
	Giturl value.Value[string]
}

func (b *baseValues) validate() error {
	if !b.Giturl.IsValid() {
		return fmt.Errorf("giturl is required")
	}
	return nil
}

func (b *baseValues) String() string {
	return fmt.Sprintf("{Giturl: %s}", b.Giturl)
}

func newBaseValues(viper *viper.Viper) (*baseValues, error) {
	values := &baseValues{
		Giturl: value.NewStringValue("giturl", "", true, viper),
	}
	if err := values.validate(); err != nil {
		return nil, fmt.Errorf("[BaseValues] error validating config:\n%v", err)
	}
	return values, nil
}
