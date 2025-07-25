package workers

// EmailWorker implements the Worker interface for email notifications
type EmailWorker struct{}

// DoWork sends an email notification and returns a status message
func (r *EmailWorker) DoWork() string {
	return "EmailWorker sent"
}
