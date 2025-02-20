package collections

import "cmp"

func CopyMap[K cmp.Ordered, V any](src map[K]V) map[K]V {
	if src == nil {
		return nil
	}
	dest := make(map[K]V, len(src))
	for k, v := range src {
		dest[k] = v
	}
	return dest
}
