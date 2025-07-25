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
func NewNotificationService(in di.NotificationServiceIn) *NotificationService {
	return &NotificationService{
		emailWorker: in.EmailWorker,
		smsWorker:   in.SMSWorker,
	}
}

// NotifyAll sends notifications through all available channels
func (r *NotificationService) NotifyAll() {
	fmt.Println(r.emailWorker.DoWork())
	fmt.Println(r.smsWorker.DoWork())
}
