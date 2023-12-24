package option

// Option is a generic struct that can hold any type of value or none.
type Option[T any] struct {
	value T
	none  bool
}

// Some returns an Option with a value.
func Some[T any](val T) Option[T] {
	return Option[T]{value: val, none: false}
}

// SomeOrNone returns an Option with a value if it exists, or None if it doesn't.
func SomeOrNone[T any](val *T) Option[T] {
	if val == nil {
		return None[T]()
	}

	return Some(*val)
}

// None returns an Option with no value.
func None[T any]() Option[T] {
	return Option[T]{value: *new(T), none: true}
}

// IsSome checks if the Option has a value.
func (o *Option[T]) IsSome() bool {
	return !o.none
}

// IsNone checks if the Option has no value.
func (o *Option[T]) IsNone() bool {
	return o.none
}

// Unwrap returns the value of the Option and a boolean indicating if it has a value.
func (o *Option[T]) Unwrap() (T, bool) {
	if o.IsNone() {
		return o.value, false
	}

	return o.value, true
}

// UnwrapOr returns the value of the Option if it exists, or a default value if it doesn't.
func (o *Option[T]) UnwrapOr(defaultVal T) T {
	if o.IsNone() {
		return defaultVal
	}

	return o.value
}
