package di

import (
	"testing"
)

// TestProvideEmailWorker tests the ProvideEmailWorker function
func TestProvideEmailWorker(t *testing.T) {
	emailWorker := ProvideEmailWorker()

	// Test EmailWorker
	emailResult := emailWorker.DoWork()
	if emailResult != "EmailWorker sent" {
		t.Errorf("ProvideEmailWorker().DoWork() = %s; want %s", emailResult, "EmailWorker sent")
	}
}

// TestProvideSMSWorker tests the ProvideSMSWorker function
func TestProvideSMSWorker(t *testing.T) {
	smsWorker := ProvideSMSWorker()

	// Test SMSWorker
	smsResult := smsWorker.DoWork()
	if smsResult != "SMSWorker sent" {
		t.Errorf("ProvideSMSWorker().DoWork() = %s; want %s", smsResult, "SMSWorker sent")
	}
}
