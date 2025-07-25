package di

import (
	"awesomeProject19/workers"
	"fmt"
	"go.uber.org/dig"
	"io"
)

// In is an alias for dig.In for better readability
type In = dig.In

// Out is an alias for dig.Out for better readability
type Out = dig.Out

type ProvideOption = dig.ProvideOption

// Container is a custom container that wraps dig.Container
type Container struct {
	container *dig.Container
}

// Provide is a custom implementation that panics on error instead of returning it
func (r *Container) Provide(constructor interface{}, opts ...ProvideOption) {
	err := r.container.Provide(constructor, opts...)
	if err != nil {
		panic(fmt.Sprintf("di.Provide error: %v", err))
	}
}

// Invoke passes through to the underlying container's Invoke method
func (r *Container) Invoke(function interface{}) error {
	return r.container.Invoke(function)
}

func Named(name string) ProvideOption {
	return dig.Name(name)
}

// New creates a new custom Container
func New() *Container {
	return &Container{
		container: dig.New(),
	}
}

// Visualize is a helper function that works with our custom Container
func Visualize(c *Container, w io.Writer) error {
	return dig.Visualize(c.container, w)
}

// NotificationServiceParams is an input struct for consuming workers with named values
type NotificationServiceParams struct {
	In

	Email     workers.Worker `name:"email_worker"`
	SMSWorker workers.Worker `name:"sms_worker"`
}

// ProvideEmailWorker creates and provides an email worker implementation
func ProvideEmailWorker() workers.Worker {
	return workers.NewEmailWorker()
}

// ProvideSMSWorker creates and provides an SMS worker implementation
func ProvideSMSWorker() workers.Worker {
	return workers.NewSMSWorker()
}
