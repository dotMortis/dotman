package value

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type StringValue struct {
	key      string
	fallback string
	required bool
	viper    *viper.Viper
}

func (v *StringValue) Value() string {
	if value := v.viper.GetString(v.key); value != "" {
		return value
	}
	return v.fallback
}

func (v *StringValue) Required() bool {
	return v.required
}

func (v *StringValue) Set(value string) error {
	value = strings.TrimSpace(value)
	if value == "" && v.required {
		return fmt.Errorf("value is required")
	}
	v.viper.Set(v.key, value)
	return nil
}

func (v *StringValue) Key() string {
	return v.key
}

func (v *StringValue) IsValid() bool {
	return valueValidate(v) == nil
}

func (v *StringValue) String() string {
	return v.Value()
}

func NewStringValue(key string, fallback string, required bool, viper *viper.Viper) *StringValue {
	return &StringValue{
		key:      key,
		fallback: fallback,
		required: required,
		viper:    viper,
	}
}
