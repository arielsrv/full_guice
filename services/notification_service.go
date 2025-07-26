package services

import (
	"fmt"
	"sync"

	"awesomeProject19/workers"
)

// NotificationService Servicio principal.
type NotificationService struct {
	workers []workers.Worker
}

// NewNotificationService creates a new notification service with the provided workers.
func NewNotificationService(params NotificationServiceParams) *NotificationService {
	return &NotificationService{
		workers: []workers.Worker{
			params.EmailWorker,
			params.SMSWorker,
		},
	}
}

// NotifyAll sends notifications through all available channels.
func (r *NotificationService) NotifyAll() {
	var wg sync.WaitGroup
	wg.Add(len(r.workers))
	for i := range r.workers {
		worker := r.workers[i]
		go func() {
			defer wg.Done()
			fmt.Println(worker.DoWork())
		}()
	}
	wg.Wait()
}
