package workers

import (
	"testing"
)

// TestEmailWorkerDoWork tests the EmailWorker implementation
func TestEmailWorkerDoWork(t *testing.T) {
	worker := &EmailWorker{}
	result := worker.DoWork()

	expected := "EmailWorker sent"
	if result != expected {
		t.Errorf("EmailWorker.DoWork() = %s; want %s", result, expected)
	}
}

// TestSMSWorkerDoWork tests the SMSWorker implementation
func TestSMSWorkerDoWork(t *testing.T) {
	worker := &SMSWorker{}
	result := worker.DoWork()

	expected := "SMSWorker sent"
	if result != expected {
		t.Errorf("SMSWorker.DoWork() = %s; want %s", result, expected)
	}
}

// TestWorkerInterface tests that both implementations satisfy the Worker interface
func TestWorkerInterface(t *testing.T) {
	var _ Worker = &EmailWorker{}
	var _ Worker = &SMSWorker{}
}
