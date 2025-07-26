package workers_test

import (
	"testing"

	"awesomeProject19/workers"

	"github.com/stretchr/testify/assert"
)

func TestNewEmailWorker(t *testing.T) {
	// Act
	worker := workers.NewEmailWorker()

	// Assert
	assert.NotNil(t, worker, "NewEmailWorker should return a non-nil worker")
	assert.IsType(t, &workers.EmailWorker{}, worker, "NewEmailWorker should return an EmailWorker instance")
}

func TestEmailWorker_DoWork(t *testing.T) {
	// Arrange
	worker := workers.NewEmailWorker()

	// Act
	result := worker.DoWork()

	// Assert
	assert.Equal(t, "EmailWorker sent", result, "DoWork should return the expected message")
}
