package registry_test

import (
	"slices"
	"testing"

	reg "github.com/bytestandard/bsugo/registry"
)

type item interface {
	MyName() string
}
type itemImpl struct {
	name string
}

func (i *itemImpl) MyName() string {
	return i.name
}
func TestRegistery(t *testing.T) {

	r := reg.NewRegistry[string]()
	r.Register("a", "aItem")
	r.Register("b", "bItem")
	r.Register("c", "cItem")
	if len(r.List()) != 3 {
		t.Errorf("Expected 3, got %d", len(r.List()))
	}
	if _, ok := r.Lookup("a"); !ok {
		t.Errorf("Expected to find a")
	}
	if _, ok := r.Lookup("d"); ok {
		t.Errorf("Expected not to find d")
	}

	names := r.List()
	if len(names) != r.Len() || !slices.Contains(names, "a") || !slices.Contains(names, "b") || !slices.Contains(names, "c") {
		t.Errorf("Expected to find a, b, and c")
	}
	r.Unregister("a")
	if len(r.List()) != 2 {
		t.Errorf("Expected 2, got %d", len(r.List()))
	}

}

func TestRegisterNilPanic(t *testing.T) {

	r2 := reg.NewRegistry[item]()
	r2.Register("a", &itemImpl{name: "a"})
	r2.Register("b", &itemImpl{name: "b"})
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	r2.Register("c", nil)
}

func TestRegisterDupPanic(t *testing.T) {

	r2 := reg.NewRegistry[item]()
	r2.Register("a", &itemImpl{name: "a"})
	r2.Register("b", &itemImpl{name: "b"})
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	r2.Register("b", &itemImpl{name: "b"})
}
