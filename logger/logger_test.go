package logger

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Log.SetLevel(WARN)
	Log.SetPrefix("test prefix")
	Log.Info("info")
	Log.Debug("debug")
	Log.Errorf("error")
	Log.Warn("warn")

	Error("error")
	Debug("debug")
	Warn("warn")
}
