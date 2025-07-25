package services

import (
	"awesomeProject19/di"
	"awesomeProject19/workers"
	"fmt"
)

// NotificationService Servicio principal
type NotificationService struct {
	emailWorker workers.Worker
	smsWorker   workers.Worker
}

// NewNotificationService creates a new notification service with the provided workers
func NewNotificationService(serviceIn di.NotificationServiceIn) *NotificationService {
	return &NotificationService{
		emailWorker: serviceIn.Email,
		smsWorker:   serviceIn.SMSWorker,
	}
}

// NotifyAll sends notifications through all available channels
func (n *NotificationService) NotifyAll() {
	fmt.Println(n.emailWorker.DoWork())
	fmt.Println(n.smsWorker.DoWork())
}
