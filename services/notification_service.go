package services

import (
	"fmt"
	"iter"
	"slices"
	"sync"

	"awesomeProject19/workers"
)

// NotificationService Servicio principal.
type NotificationService struct {
	workers iter.Seq[workers.Worker]
}

// NewNotificationService creates a new notification service with the provided workers.
func NewNotificationService(params NotificationServiceParams) *NotificationService {
	return &NotificationService{
		workers: slices.Values([]workers.Worker{
			params.EmailWorker,
			params.SMSWorker,
		}),
	}
}

// NotifyAll sends notifications through all available channels.
func (r *NotificationService) NotifyAll() {
	var wg sync.WaitGroup
	for worker := range r.workers {
		wg.Go(func() {
			fmt.Println(worker.DoWork())
		})
	}
	wg.Wait()
}
