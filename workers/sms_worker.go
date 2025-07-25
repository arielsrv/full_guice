package workers

// SMSWorker implements the Worker interface for SMSWorker notifications
type SMSWorker struct{}

// NewSMSWorker creates a new instance of SMSWorker
func NewSMSWorker() Worker {
	return &SMSWorker{}
}

// DoWork sends an SMSWorker notification and returns a status message
func (r *SMSWorker) DoWork() string {
	return "SMSWorker sent"
}
