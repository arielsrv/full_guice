package main

import (
	"awesomeProject19/di"
	"awesomeProject19/services"
	"awesomeProject19/workers"
)

func main() {
	container := di.New()

	// Proveer implementaciones
	container.Provide(func() workers.Worker { return workers.NewEmailWorker() }, di.Named("email_worker"))
	container.Provide(func() workers.Worker { return workers.NewSMSWorker() }, di.Named("sms_worker"))

	container.Provide(services.NewNotificationService)

	// Ejecutar
	err := container.Invoke(func(service *services.NotificationService) {
		service.NotifyAll()
	})
	if err != nil {
		panic(err)
	}
}
