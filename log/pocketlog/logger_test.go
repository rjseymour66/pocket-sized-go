package pocketlog_test

import (
	"os"

	"github.com/rjseymour66/pocket-sized-go/tree/master/log/pocketlog"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, os.Stdout)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}
