package workers

// SMSWorker implements the Worker interface for SMSWorker notifications
type SMSWorker struct{}

// DoWork sends an SMSWorker notification and returns a status message
func (r *SMSWorker) DoWork() string {
	return "SMSWorker sent"
}
