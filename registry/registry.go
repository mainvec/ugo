package registry

import "sync"

// Registry is a generic data structure that maps strings to values of any type.
type Registry[T any] struct {
	m  map[string]T
	mu sync.RWMutex
}

// NewRegistry creates a new instance of the Registry type.
// The Registry is a generic data structure that maps strings to values of any type.
func NewRegistry[T any]() *Registry[T] {
	return &Registry[T]{
		m: map[string]T{},
	}
}

func isNil(t any) bool {
	return t == nil
}

// Register adds a new value to the registry.
// it panics if the name is empty, the value is nil it, or if the name is already in use.
func (r *Registry[T]) Register(name string, t T) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if len(name) == 0 {
		panic("Registering with empty name")
	}
	if isNil(t) {
		panic("Registering nil value")
	}

	if _, dup := r.m[name]; dup {
		panic("Duplicated registration")
	}

	r.m[name] = t
}

// Lookup returns the value associated with the given name.
// The second return value indicates whether the name was found in the registry.
func (r *Registry[T]) Lookup(name string) (T, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.m[name]
	return t, ok
}

// Unregister removes the value associated with the given name.
func (r *Registry[T]) Unregister(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.m, name)
}

// List returns a list of all the names in the registry.
func (r *Registry[T]) List() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var list []string
	for k := range r.m {
		list = append(list, k)
	}
	return list
}

func (r *Registry[T]) ListValues() []T {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]T, 0, len(r.m))
	for _, v := range r.m {
		list = append(list, v)
	}
	return list
}

func (r *Registry[T]) Len() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.m)
}
