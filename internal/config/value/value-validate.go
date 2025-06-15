package value

import "fmt"

func valueValidate[T string | int | bool | float64](value Value[T]) error {
	if value.Required() && !isSet(value.Value()) {
		return fmt.Errorf("value is required")
	}
	return nil
}

func isSet[T string | int | bool | float64](value T) bool {
	switch v := any(value).(type) {
	case string:
		return v != ""
	default:
		return true
	}
}
