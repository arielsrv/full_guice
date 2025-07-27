package services

import (
	"fmt"
	"iter"
	"slices"
	"sync"

	"awesomeProject19/workers"
)

type INotificationService interface {
	NotifyAll()
}

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
		wg.Add(1)
		go func(w workers.Worker) {
			defer wg.Done()
			fmt.Println(w.DoWork())
		}(worker)
	}
	wg.Wait()
}
