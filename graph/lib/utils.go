package lib

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Head[T any](ts []T) *T {
	if len(ts) == 0 {
		return nil
	}
	return &ts[0]
}
