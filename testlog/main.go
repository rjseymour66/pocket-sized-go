package main

import (
	"fmt"
	"os"

	"github.com/rjseymour66/pocket-sized-go/tree/master/log/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, os.Stdout)

	fmt.Println("Hello world")
}
