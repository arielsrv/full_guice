package services

import (
	"awesomeProject19/di"
	"awesomeProject19/workers"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// TestNewNotificationService tests the creation of NotificationService
func TestNewNotificationService(t *testing.T) {
	// Create mock workers
	emailWorker := &workers.EmailWorker{}
	smsWorker := &workers.SMSWorker{}

	// Create input struct
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
}

// TestNotificationServiceNotifyAll tests the NotifyAll method
func TestNotificationServiceNotifyAll(t *testing.T) {
	// Create mock workers
	emailWorker := &workers.EmailWorker{}
	smsWorker := &workers.SMSWorker{}

	// Create input struct
	in := di.NotificationServiceParams{
		Email:     emailWorker,
		SMSWorker: smsWorker,
	}

	// Create service
	service := NewNotificationService(in)

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call NotifyAll
	service.NotifyAll()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Verify output contains expected messages
	if !strings.Contains(output, "EmailWorker sent") {
		t.Error("NotifyAll() did not print 'EmailWorker sent'")
	}

	if !strings.Contains(output, "SMSWorker sent") {
		t.Error("NotifyAll() did not print 'SMSWorker sent'")
	}
}
