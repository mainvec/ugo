package collections

import "cmp"

// ArraysEqualvant compare two arrays and return true if they are
// have same length and same elements, regardless of the order of the elements. Otherwise, return false.
func ArraysEqualvant[V cmp.Ordered](a1 []V, a2 []V) bool {
	if len(a1) != len(a2) {
		return false
	}
	for _, v := range a1 {
		if !ArraysContains(a2, v) {
			return false
		}
	}
	return true
}

//Arrays contians

func ArraysContains[V cmp.Ordered](arr []V, value V) bool {
	for _, v := range arr {

		if v == value {
			return true
		}
	}
	return false
}
