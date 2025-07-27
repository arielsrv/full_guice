package workers_test

import (
	"testing"

	"awesomeProject19/workers"

	"github.com/stretchr/testify/assert"
)

func TestNewSMSWorker(t *testing.T) {
	// Act
	worker := workers.NewSMSWorker()

	// Assert
	assert.NotNil(t, worker, "NewSMSWorker should return a non-nil worker")
	assert.IsType(t, &workers.SMSWorker{}, worker, "NewSMSWorker should return an SMSWorker instance")
}

func TestSMSWorker_DoWork(t *testing.T) {
	// Arrange
	worker := workers.NewSMSWorker()

	// Act
	result := worker.DoWork()

	// Assert
	assert.Equal(t, "SMSWorker sent", result, "DoWork should return the expected message")
}
