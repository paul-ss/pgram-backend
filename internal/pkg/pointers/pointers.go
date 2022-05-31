package pointers

func New[T constraint](a T) *T {
	return &a
}

type constraint interface {
	int | int8 | int16 | int32 | int64 | uint |
		uint8 | uint16 | uint32 | uint64 | float32 | float64 |
		string
}
