package main

import (
	"log/pocketlog"
	"time"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo)

	lgr.Infof("A little copying is better than a little dependency.")
	lgr.Errorf("Errors are values. Documentation is for %s", "users")
	lgr.Debugf("Make the zero (%d) value useful.", 0)

	lgr.Infof("Hallo, %d %v", 2022, time.Now())
}
