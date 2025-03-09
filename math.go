package utils

type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

func Abs[T SignedInteger](x T) T {
	if x < 0 {
		return x * -1
	}
	return x
}
