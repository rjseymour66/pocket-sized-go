package main

import (
	"bufio"
	"fmt"
	"gordle/gordle"
	"os"
)

const maxAttempts = 6

func main() {

	corpus, err := gordle.ReadCorpus("corpus/english.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to read corpus: %s", err)
	}

	// Create the game
	g, err := gordle.New(bufio.NewReader(os.Stdin), corpus, maxAttempts)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to start game: %s", err)
		return
	}

	// Run the game
	g.Play()
}
