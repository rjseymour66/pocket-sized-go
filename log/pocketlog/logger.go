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
	format = fmt.Sprintf("DEBUG:\t%s", format)

	l.logf(l.output, format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	// validate it is the correct logger
	if l.threshold > LevelInfo {
		return
	}

	format = fmt.Sprintf("INFO:\t%s", format)

	l.logf(l.output, format, args...)
}

// Errorf formats and prints a message if the log level is error.
func (l *Logger) Errorf(format string, args ...any) {
	format = fmt.Sprintf("ERROR:\t%s", format)
	l.logf(l.output, format, args...)
}

// logf prints the message to the output.
// Add decorations here, if any.
func (l *Logger) logf(output io.Writer, format string, args ...any) {
	_, _ = fmt.Fprintf(output, format+"\n", args...)
}

type jsonLog struct {
	Level   Level  `json:"level"`
	Message string `json:"message"`
}
