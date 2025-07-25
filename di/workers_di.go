package di

import (
	"awesomeProject19/workers"
	"go.uber.org/dig"
)

// In is an alias for dig.In for better readability
type In = dig.In

// Out is an alias for dig.Out for better readability
type Out = dig.Out

type Container = dig.Container

func New() *Container {
	return dig.New()
}

// NotificationServiceIn is an input struct for consuming workers with named values
type NotificationServiceIn struct {
	In

	Email     workers.Worker `name:"email_worker"`
	SMSWorker workers.Worker `name:"sms_worker"`
}

// Workers is an output struct for providing workers with named values
type Workers struct {
	Out

	EmailWorker workers.Worker `name:"email_worker"`
	SMSWorker   workers.Worker `name:"sms_worker"`
}

// ProvideWorkers creates and provides worker implementations
func ProvideWorkers() Workers {
	return Workers{
		EmailWorker: &workers.EmailWorker{},
		SMSWorker:   &workers.SMSWorker{},
	}
}
