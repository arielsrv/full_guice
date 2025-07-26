package workers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmailWorker(t *testing.T) {
	// Act
	worker := NewEmailWorker()

	// Assert
	assert.NotNil(t, worker, "NewEmailWorker should return a non-nil worker")
	assert.IsType(t, &EmailWorker{}, worker, "NewEmailWorker should return an EmailWorker instance")
}

func TestEmailWorker_DoWork(t *testing.T) {
	// Arrange
	worker := NewEmailWorker()

	// Act
	result := worker.DoWork()

	// Assert
	assert.Equal(t, "EmailWorker sent", result, "DoWork should return the expected message")
}
