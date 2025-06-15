package value

import "fmt"

type Value[T string | int | bool | float64] interface {
	fmt.Stringer
	Value() T
	Required() bool
	Set(T) error
	Key() string
	IsValid() bool
}
