package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns a logger, ready to log at the required threshold.
func New(threshold Level, opts ...Option) *Logger {
	// set defaults
	lgr := &Logger{threshold: threshold, output: os.Stdout}

	// add config options
	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

// Debugf formats and prints a message if the log level is debug or higher.
// The default output is Stdout.
func (l *Logger) Debugf(format string, args ...any) {
	// validate it is the correct logger
	if l.threshold > LevelDebug {
		return
	}
	// making sure we can safely write to the output
	if l.output == nil {
		l.output = os.Stdout
	}
	l.logf(l.output, format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	// validate it is the correct logger
	if l.threshold > LevelInfo {
		return
	}
	// making sure we can safely write to the output
	if l.output == nil {
		l.output = os.Stdout
	}

	l.logf(l.output, format, args...)
}

// Errorf formats and prints a message if the log level is error.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold < LevelError {
		return
	}
	// making sure we can safely write to the output
	if l.output == nil {
		l.output = os.Stdout
	}

	l.logf(l.output, format, args...)
}

// logf prints the message to the output.
// Add decorations here, if any.
func (l *Logger) logf(output io.Writer, format string, args ...any) {
	_, _ = fmt.Fprintf(output, format+"\n", args...)
}
