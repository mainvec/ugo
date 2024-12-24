package collections

import "golang.org/x/exp/constraints"

func CopyMap[K constraints.Ordered, V any](src map[K]V) map[K]V {
	if src == nil {
		return nil
	}
	dest := make(map[K]V, len(src))
	for k, v := range src {
		dest[k] = v
	}
	return dest
}
