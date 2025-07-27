package di

import (
	"fmt"

	"go.uber.org/dig"
)

// In is an alias for dig.In for better readability.
type In = dig.In

// Out is an alias for dig.Out for better readability.
type Out = dig.Out

type ProvideOption = dig.ProvideOption

// Registry is a custom container that wraps dig.Container.
type Registry struct {
	container *dig.Container
}

// New creates a new Registry instance with a fresh dig container.
func New() *Registry {
	return &Registry{container: dig.New()}
}

// Provide is a custom implementation that panics on error instead of returning it.
func (r *Registry) Provide(constructor interface{}, opts ...ProvideOption) {
	err := r.container.Provide(constructor, opts...)
	if err != nil {
		panic(fmt.Sprintf("di.Provide error: %v", err))
	}
}

// Invoke passes through to the underlying container's Invoke method.
func (r *Registry) Invoke(function interface{}) error {
	return r.container.Invoke(function)
}

// GetInstance retrieves a dependency of type T from the container using generics.
// This provides a type-safe alternative to Invoke for simple dependency retrieval.
func GetInstance[T any]() T {
	var result T
	err := Container.Invoke(func(dep T) {
		result = dep
	})
	if err != nil {
		panic(fmt.Sprintf("di.GetInstance error: %v", err))
	}
	return result
}

// GetNamed retrieves a named dependency of type T from the container using generics.
// For now, this is a placeholder that will be implemented using a different approach.
// Named dependencies in dig require struct tags, which can't be dynamically created.
// Users should use the regular GetInstance[T]() for non-named dependencies or create specific
// getter functions for named dependencies.
func GetNamed[T any](container *Registry, name string) T {
	panic("GetNamed is not yet implemented. Use specific getter functions for named dependencies or regular GetInstance[T]() for non-named ones.")
}

func Named(name string) ProvideOption {
	return dig.Name(name)
}

func As(constructor interface{}) ProvideOption {
	return dig.As(constructor)
}

var Container = &Registry{container: dig.New()}

func Reset() {
	Container = &Registry{container: dig.New()}
}
