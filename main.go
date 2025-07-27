package main

import (
	"awesomeProject19/di"
	"awesomeProject19/services"
	"awesomeProject19/workers"
)

func main() {
	// Proveer implementaciones
	di.Container.Provide(func() workers.Worker { return workers.NewEmailWorker() }, di.Named("email_worker"))
	di.Container.Provide(func() workers.Worker { return workers.NewSMSWorker() }, di.Named("sms_worker"))
	di.Container.Provide(services.NewNotificationService, di.As(new(services.INotificationService)))

	// Ejecutar usando la nueva API genérica
	notificationService := di.GetInstance[services.INotificationService]()
	notificationService.NotifyAll()
}
