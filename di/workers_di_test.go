package di

import (
	"testing"
)

// TestProvideWorkers tests the ProvideWorkers function
func TestProvideWorkers(t *testing.T) {
	workers := ProvideWorkers()

	// Test EmailWorker worker
	emailResult := workers.EmailWorker.DoWork()
	if emailResult != "EmailWorker sent" {
		t.Errorf("ProvideWorkers().EmailWorker.DoWork() = %s; want %s", emailResult, "EmailWorker sent")
	}

	// Test SMSWorker worker
	smsResult := workers.SMSWorker.DoWork()
	if smsResult != "SMSWorker sent" {
		t.Errorf("ProvideWorkers().SMSWorker.DoWork() = %s; want %s", smsResult, "SMSWorker sent")
	}
}
