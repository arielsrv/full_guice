package main

import (
	"awesomeProject19/di"
	"awesomeProject19/services"
	"fmt"
	"os"
)

func main() {
	container := di.New()

	// Proveer implementaciones
	container.Provide(di.ProvideEmailWorker, di.Named("email_worker"))
	container.Provide(di.ProvideSMSWorker, di.Named("sms_worker"))
	container.Provide(services.NewNotificationService)

	// Mostrar el grafo de dependencias
	fmt.Println("Grafo de dependencias:")
	if err := di.Visualize(container, os.Stdout); err != nil {
		fmt.Printf("Error al visualizar el grafo: %v\n", err)
	}
	fmt.Println("\n---")

	// Ejecutar
	err := container.Invoke(func(service *services.NotificationService) {
		service.NotifyAll()
	})
	if err != nil {
		panic(err)
	}
}
