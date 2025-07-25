package main

import (
	"awesomeProject19/di"
	"awesomeProject19/services"
)

func main() {
	container := di.New()

	// Proveer implementaciones
	container.Provide(di.ProvideEmailWorker, di.Named("email_worker"))
	container.Provide(di.ProvideSMSWorker, di.Named("sms_worker"))
	container.Provide(services.NewNotificationService)

	// Ejecutar
	err := container.Invoke(func(service *services.NotificationService) {
		service.NotifyAll()
	})
	if err != nil {
		panic(err)
	}
}
