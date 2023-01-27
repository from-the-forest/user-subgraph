package lib

import "reflect"

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

// Merge two scructs of the same type together
func Merge(a, b interface{}) interface{} {
	val := reflect.ValueOf(a).Elem()
	val2 := reflect.ValueOf(b).Elem()
	for i := 0; i < val.NumField(); i++ {
		val.Field(i).Set(val2.Field(i))
	}
	return val
}
