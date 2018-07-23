package config

import (
	"testing"
	"github.com/op/go-logging"
)

func TestDebugOffWhenTesting(t *testing.T) {
	if Logger.IsEnabledFor(logging.DEBUG) {
		t.Errorf("Logging should be info")
	}
}
