package collections_test

import (
	"testing"

	"github.com/workoak/wogo/collections"
)

func TestArrayContains(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !collections.ArraysContains(arr, 3) {
		t.Fatal("wanted 3 in arr")
	}
	if collections.ArraysContains(arr, 6) {
		t.Fatal("wanted 6 not in arr")
	}
}

func TestArrayEqualvant(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 4, 3, 5}
	if !collections.ArraysEqualvant(arr1, arr2) {
		t.Fatal("wanted arr1 == arr2")
	}
	arr2 = []int{1, 2, 3, 4}
	if collections.ArraysEqualvant(arr1, arr2) {
		t.Fatal("wanted arr1 != arr2")
	}
}
