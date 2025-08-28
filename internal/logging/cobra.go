package logging

import (
	"log/slog"
	"strings"
)

// CobraLogger is a wrapper around slog to use it as the default output for cobra output.
type CobraLogger struct {
	logger *slog.Logger
}

// NewCobraLogger returns a wrapper for logging cobra output.
func NewCobraLogger(logger *slog.Logger) *CobraLogger {
	return &CobraLogger{
		logger: logger,
	}
}

// Write prints any message with slog.logger on WARN level.
func (c *CobraLogger) Write(p []byte) (n int, err error) {
	// Cobra can add \n or spaces to the end of the line, so remove it for pretty printing.
	msg := strings.TrimSpace(string(p))
	c.logger.Warn(msg)

	return len(p), nil
}
