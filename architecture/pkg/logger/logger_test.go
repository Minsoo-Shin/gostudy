package logger

import (
	"testing"
)

func TestA(t *testing.T) {
	logger := New()
	logger.Info("Hello, world!",
		String(ID, "12312515"))
}
