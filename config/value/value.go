package value

type Value[T string | int | bool | float64] interface {
	Value() T
	Required() bool
	Set(T) error
	Key() string
	IsValid() bool
}
