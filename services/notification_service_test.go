package services

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"awesomeProject19/workers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	mockworkers "awesomeProject19/mocks/workers"
)

func TestNewNotificationService(t *testing.T) {
	// Arrange
	mockT := new(testing.T)
	emailWorker := mockworkers.NewMockWorker(mockT)
	smsWorker := mockworkers.NewMockWorker(mockT)

	params := NotificationServiceParams{
		EmailWorker: emailWorker,
		SMSWorker:   smsWorker,
	}

	// Act
	service := NewNotificationService(params)

	// Assert
	assert.NotNil(t, service, "NewNotificationService should return a non-nil service")
	assert.Len(t, service.workers, 2, "Service should have 2 workers")
	assert.Equal(t, emailWorker, service.workers[0], "First worker should be the email worker")
	assert.Equal(t, smsWorker, service.workers[1], "Second worker should be the SMS worker")
}

func TestNotificationService_NotifyAll(t *testing.T) {
	// Arrange
	mockT := new(testing.T)
	emailWorker := mockworkers.NewMockWorker(mockT)
	smsWorker := mockworkers.NewMockWorker(mockT)

	emailWorker.EXPECT().DoWork().Return("Email notification sent")
	smsWorker.EXPECT().DoWork().Return("SMS notification sent")

	params := NotificationServiceParams{
		EmailWorker: emailWorker,
		SMSWorker:   smsWorker,
	}

	service := NewNotificationService(params)

	// Capture stdout to verify output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Act
	service.NotifyAll()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	require.NoError(t, err, "Failed to capture stdout")

	// Assert
	output := buf.String()
	assert.Contains(t, output, "Email notification sent", "Output should contain email notification message")
	assert.Contains(t, output, "SMS notification sent", "Output should contain SMS notification message")
}

func TestNotificationService_NotifyAll_EmptyWorkers(t *testing.T) {
	// Arrange
	service := &NotificationService{
		workers: []workers.Worker{},
	}

	// Capture stdout to verify no output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Act
	service.NotifyAll()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	require.NoError(t, err, "Failed to capture stdout")

	// Assert
	output := buf.String()
	assert.Empty(t, strings.TrimSpace(output), "Output should be empty when there are no workers")
}
