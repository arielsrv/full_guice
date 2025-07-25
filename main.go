package main

import (
	"awesomeProject19/di"
	"awesomeProject19/services"
)

func main() {
	container := di.New()

	// Proveer implementaciones
	container.Provide(di.ProvideWorkers)
	container.Provide(services.NewNotificationService)

	// Ejecutar
	err := container.Invoke(func(service *services.NotificationService) {
		service.NotifyAll()
	})
	if err != nil {
		panic(err)
	}
}
