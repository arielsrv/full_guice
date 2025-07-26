package services

import (
	"awesomeProject19/di"
	"awesomeProject19/workers"
)

// NotificationServiceParams is an input struct for consuming workers with named values.
type NotificationServiceParams struct {
	di.In

	EmailWorker workers.Worker `name:"email_worker"`
	SMSWorker   workers.Worker `name:"sms_worker"`
}
