package services

import (
	"awesomeProject19/di"
	"awesomeProject19/workers"
	"testing"
)

// MockWorker is a mock implementation of the workers.Worker interface for testing
type MockWorker struct {
	DoWorkFunc func() string
}

func (m *MockWorker) DoWork() string {
	if m.DoWorkFunc != nil {
		return m.DoWorkFunc()
	}
	return "mock response"
}

// TestNotificationServiceWithMocks tests the NotificationService with mock workers
func TestNotificationServiceWithMocks(t *testing.T) {
	// Create mock workers with custom behavior
	emailWorker := &MockWorker{
		DoWorkFunc: func() string {
			return "Custom emailWorker response"
		},
	}

	smsWorker := &MockWorker{
		DoWorkFunc: func() string {
			return "Custom SMSWorker response"
		},
	}

	// Create input struct with mocks
	in := di.NotificationServiceParams{
		Email:     emailWorker,
		SMSWorker: smsWorker,
	}

	// Create service
	service := NewNotificationService(in)

	// Verify service has the correct workers
	if service.emailWorker != emailWorker {
		t.Error("NewNotificationService did not set the emailWorker worker correctly")
	}

	if service.smsWorker != smsWorker {
		t.Error("NewNotificationService did not set the smsWorker worker correctly")
	}

	// Test that the service uses the mock implementations correctly
	// This is an indirect test since NotifyAll prints to stdout
	// In a real-world scenario, we might refactor NotifyAll to return values instead
}

// TestWorkerInterfaceWithMock tests that the mock implementation satisfies the Worker interface
func TestWorkerInterfaceWithMock(t *testing.T) {
	var _ workers.Worker = &MockWorker{}
}
