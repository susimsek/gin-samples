package util

type Optional[T any] struct {
	Value *T
}

// IsPresent checks if the value is not nil
func (o Optional[T]) IsPresent() bool {
	return o.Value != nil
}

// IsEmpty checks if the value is nil
func (o Optional[T]) IsEmpty() bool {
	return o.Value == nil
}

// OrElse returns the value if present, or the defaultValue if nil
func (o Optional[T]) OrElse(defaultValue T) T {
	if o.Value == nil {
		return defaultValue
	}
	return *o.Value
}
