package pocketlog_test

import (
	"log/pocketlog"
	"testing"
)

const (
	debugMessage = "Why write I still all one, ever the same,"
	infoMessage  = "And keep invention in a noted weed,"
	errorMessage = "That every word doth almost tell my name,"
)

func TestLogger_DebugfInfofErrorf(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: "DEBUG:\t" + debugMessage + "\n" + "INFO:\t" + infoMessage + "\n" + "ERROR:\t" + errorMessage + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: "INFO:\t" + infoMessage + "\n" + "ERROR:\t" + errorMessage + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: "ERROR:\t" + errorMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Errorf(errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

// testWriter is a struct that implements io.Writer
type testWriter struct {
	contents string
}

// Write implements the io.Writer interface
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
