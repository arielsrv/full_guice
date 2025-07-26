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

// Container is a custom container that wraps dig.Container.
type Container struct {
	container *dig.Container
}

// Provide is a custom implementation that panics on error instead of returning it.
func (r *Container) Provide(constructor interface{}, opts ...ProvideOption) {
	err := r.container.Provide(constructor, opts...)
	if err != nil {
		panic(fmt.Sprintf("di.Provide error: %v", err))
	}
}

// Invoke passes through to the underlying container's Invoke method.
func (r *Container) Invoke(function interface{}) error {
	return r.container.Invoke(function, dig.FillInvokeInfo(nil))
}

func Named(name string) ProvideOption {
	return dig.Name(name)
}

// New creates a new custom Container.
func New() *Container {
	return &Container{
		container: dig.New(),
	}
}
