package util

type Optional[T any] struct {
	Value *T
}

func (o Optional[T]) IsPresent() bool {
	return o.Value != nil
}

func (o Optional[T]) OrElse(defaultValue T) T {
	if o.Value == nil {
		return defaultValue
	}
	return *o.Value
}
